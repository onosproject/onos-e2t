// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/kubernetes"
	"github.com/onosproject/helmit/pkg/simulation"
	"github.com/onosproject/helmit/pkg/util/async"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
)

const controlPort = 5000

var log = logging.GetLogger("sim", "e2")

// SimSuite is the primary onos-e2t simulation suite
type SimSuite struct {
	simulation.Suite
	apps []*simApp
}

// SetupSimulation :: simulation
func (s *SimSuite) SetupSimulation(sim *simulation.Simulator) error {
	err := helm.Chart("sd-ran").
		Release("sd-ran").
		Set("import.onos-config.enabled", false).
		Set("import.onos-e2sub.enabled", true).
		Set("global.storage.consensus.enabled", "true").
		Set("onos-topo.image.tag", "latest").
		Set("onos-e2t.image.tag", "latest").
		Set("onos-e2sub.image.tag", "latest").
		Set("ran-simulator.image.tag", "latest").
		Set("global.image.registry", sim.Arg("registry").String("")).
		Set("onos-e2t.logging.loggers.root.level", "debug").
		Install(true)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// SetupSimulator :: simulation
func (s *SimSuite) SetupSimulator(sim *simulation.Simulator) error {
	err := helm.
		Chart("ran-simulator").
		Release(sim.Name).
		Install(true)
	if err != nil {
		log.Error(err)
		return err
	}

	client, err := kubernetes.New()
	if err != nil {
		log.Error(err)
		return err
	}

	topoSdkClient, err := utils.NewTopoClient()
	if err != nil {
		return err
	}

	nodeID := utils.GetTestNodeID(nil)

	cells, err := topoSdkClient.GetCells(context.Background(), nodeID)
	if err != nil {
		return err
	}
	cellObjectID := cells[0].CellObjectID

	numApps := sim.Arg("app-count").Int(1)
	numInstances := sim.Arg("replica-count").Int(1)
	numSubs := sim.Arg("sub-count").Int(1)

	s.apps = make([]*simApp, numApps)
	for i := 0; i < numApps; i++ {
		appID := fmt.Sprintf("sim-%s-%d-%d", sim.Name, sim.Process, i+1)
		instances := make([]*simAppInstance, numInstances)
		for j := 0; j < numInstances; j++ {
			instance := &simAppInstance{
				name:    fmt.Sprintf("%s-%d", appID, j),
				address: fmt.Sprintf("%s-%d.%s:%d", appID, j, appID, controlPort),
				client:  client,
			}

			for k := 0; k < numSubs; k++ {
				sub := &simAppSub{
					name:         fmt.Sprintf("%s-%d", appID, k+1),
					instance:     instance,
					nodeID:       string(nodeID),
					cellObjectID: cellObjectID,
					reportPeriod: uint32((k + 1) * 5 * 1000),
					granularity:  500,
				}
				instance.subs = append(instance.subs, sub)
			}

			instances[j] = instance
		}

		app := &simApp{
			name:      appID,
			instances: instances,
			client:    client,
		}

		err := app.start()
		if err != nil {
			return err
		}
		s.apps[i] = app
	}
	return nil
}

// ScheduleSimulator :: simulation
func (s *SimSuite) ScheduleSimulator(sim *simulation.Simulator) {
	sim.Schedule("subscribe", s.SimulateSubscribe, 10*time.Second, 5)
	sim.Schedule("unsubscribe", s.SimulateUnsubscribe, 30*time.Second, 3)
	sim.Schedule("crash", s.SimulateCrash, 1*time.Minute, 4)
}

func (s *SimSuite) getRandInstance(predicate func(instance *simAppInstance) bool) (*simAppInstance, bool) {
	var instances []*simAppInstance
	for _, app := range s.apps {
		for _, instance := range app.instances {
			if predicate(instance) {
				instances = append(instances, instance)
			}
		}
	}

	if len(instances) == 0 {
		return nil, false
	}

	instance := instances[rand.Intn(len(instances))]
	return instance, true
}

func (s *SimSuite) getRandSub(predicate func(sub *simAppSub) bool) (*simAppSub, bool) {
	var subs []*simAppSub
	for _, app := range s.apps {
		for _, instance := range app.instances {
			if instance.ready() {
				for _, sub := range instance.subs {
					if predicate(sub) {
						subs = append(subs, sub)
						break
					}
				}
			}
		}
	}

	if len(subs) == 0 {
		return nil, false
	}

	sub := subs[rand.Intn(len(subs))]
	return sub, true
}

func (s *SimSuite) SimulateSubscribe(sim *simulation.Simulator) error {
	sub, ok := s.getRandSub(func(sub *simAppSub) bool {
		return !sub.running()
	})
	if !ok {
		return nil
	}
	return sub.start()
}

func (s *SimSuite) SimulateUnsubscribe(sim *simulation.Simulator) error {
	sub, ok := s.getRandSub(func(sub *simAppSub) bool {
		return sub.running()
	})
	if !ok {
		return nil
	}
	return sub.stop()
}

func (s *SimSuite) SimulateCrash(sim *simulation.Simulator) error {
	instance, ok := s.getRandInstance(func(instance *simAppInstance) bool {
		return instance.ready()
	})
	if !ok {
		return nil
	}
	return instance.crash()
}

type simApp struct {
	name      string
	instances []*simAppInstance
	client    kubernetes.Client
}

func (s *simApp) start() error {
	log.Infof("Starting app '%s'", s.name)
	ss := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: s.name,
			Labels: map[string]string{
				"app": s.name,
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: s.name,
			Replicas:    pointer.Int32Ptr(int32(len(s.instances))),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": s.name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": s.name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "sim-app",
							Image:           "onosproject/onos-e2t-sim-app:latest",
							ImagePullPolicy: corev1.PullIfNotPresent,
							Args: []string{
								s.name,
								"$(POD_NAME)",
							},
							Env: []corev1.EnvVar{
								{
									Name: "POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "control",
									ContainerPort: controlPort,
								},
							},
							ReadinessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									TCPSocket: &corev1.TCPSocketAction{
										Port: intstr.FromInt(controlPort),
									},
								},
								InitialDelaySeconds: 5,
								TimeoutSeconds:      10,
								FailureThreshold:    6,
							},
						},
					},
				},
			},
		},
	}

	_, err := s.client.Clientset().
		AppsV1().
		StatefulSets(s.client.Namespace()).
		Create(context.Background(), ss, metav1.CreateOptions{})
	if err != nil {
		log.Error(err)
		return err
	}

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: s.name,
			Labels: map[string]string{
				"app": s.name,
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": s.name,
			},
			ClusterIP: corev1.ClusterIPNone,
			Ports: []corev1.ServicePort{
				{
					Name: "control",
					Port: controlPort,
				},
			},
		},
	}

	_, err = s.client.Clientset().
		CoreV1().
		Services(s.client.Namespace()).
		Create(context.Background(), svc, metav1.CreateOptions{})
	if err != nil {
		log.Error(err)
		return err
	}
	return async.IterAsync(len(s.instances), func(i int) error {
		return s.instances[i].start()
	})
}

type simAppInstance struct {
	name    string
	address string
	subs    []*simAppSub
	client  kubernetes.Client
	running bool
	mu      sync.RWMutex
}

func (s *simAppInstance) ready() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.running
}

func (s *simAppInstance) start() error {
	t := time.Now()
	if err := s.awaitPodReady(); err != nil {
		log.Error(err)
		return err
	}
	go s.streamLogs(t)
	return nil
}

func (s *simAppInstance) awaitPodReady() error {
	log.Infof("Waiting for pod '%s'", s.name)
	err := backoff.Retry(func() error {
		s.mu.Lock()
		running := s.running
		s.mu.Unlock()
		if running {
			return nil
		}

		pod, err := s.client.Clientset().
			CoreV1().
			Pods(s.client.Namespace()).
			Get(context.Background(), s.name, metav1.GetOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				return err
			}
			return backoff.Permanent(err)
		}
		if pod.Status.Phase != corev1.PodRunning {
			return errors.New("retry")
		}
		if len(pod.Status.ContainerStatuses) == 0 {
			return errors.New("retry")
		}
		for _, container := range pod.Status.ContainerStatuses {
			if !container.Ready {
				return errors.New("retry")
			}
		}
		s.running = true
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return err
	}
	return nil
}

func (s *simAppInstance) streamLogs(since time.Time) {
	log.Infof("Following pod '%s'", s.name)
	t := metav1.NewTime(since)
	req := s.client.Clientset().
		CoreV1().
		Pods(s.client.Namespace()).
		GetLogs(s.name, &corev1.PodLogOptions{
			Container: "sim-app",
			Follow:    true,
			SinceTime: &t,
		})
	reader, err := req.Stream(context.Background())
	if err != nil {
		log.Error(err)
		return
	}
	defer reader.Close()

	// Stream the logs to stdout
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, scanner.Text())
	}
}

func (s *simAppInstance) crash() error {
	s.mu.Lock()
	running := s.running
	s.mu.Unlock()
	if !running {
		return nil
	}

	s.mu.Lock()
	s.running = false
	for _, sub := range s.subs {
		sub.reset()
	}
	s.mu.Unlock()

	log.Infof("Crashing pod '%s'", s.name)
	err := s.client.Clientset().
		CoreV1().
		Pods(s.client.Namespace()).
		Delete(context.Background(), s.name, metav1.DeleteOptions{})
	if err != nil {
		log.Error(err)
		return err
	}
	return s.start()
}

type simAppSub struct {
	name         string
	instance     *simAppInstance
	nodeID       string
	cellObjectID string
	reportPeriod uint32
	granularity  uint32
	started      bool
	mu           sync.RWMutex
}

func (s *simAppSub) running() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.started
}

func (s *simAppSub) start() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.started {
		return nil
	}

	log.Infof("Starting '%s' subscription '%s' on '%s'", s.nodeID, s.name, s.instance.name)
	conn, err := grpc.Dial(s.instance.address, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return err
	}
	defer conn.Close()
	client := NewSimServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &StartSubscriptionRequest{
		SubscriptionId: s.name,
		NodeId:         s.nodeID,
		CellObjectId:   s.cellObjectID,
		ReportPeriod:   s.reportPeriod,
		Granularity:    s.granularity,
	}
	_, err = client.StartSubscription(ctx, request)
	if err != nil {
		log.Error(err)
		return err
	}
	s.started = true
	return nil
}

func (s *simAppSub) stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.started {
		return nil
	}

	log.Infof("Stopping '%s' subscription '%s' on '%s'", s.nodeID, s.name, s.instance.name)
	conn, err := grpc.Dial(s.instance.address, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return err
	}
	defer conn.Close()
	client := NewSimServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &StopSubscriptionRequest{
		SubscriptionId: s.name,
		NodeId:         s.nodeID,
	}
	_, err = client.StopSubscription(ctx, request)
	if err != nil {
		log.Error(err)
		return err
	}
	s.started = false
	return nil
}

func (s *simAppSub) reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.started = false
}
