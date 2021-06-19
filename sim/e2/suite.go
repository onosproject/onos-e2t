// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"fmt"
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/kubernetes"
	"github.com/onosproject/helmit/pkg/simulation"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
	"math/rand"
	"time"
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

	objects, err := utils.GetControlRelationObjects()
	if err != nil {
		log.Error(err)
		return err
	}

	var nodeIDs []topoapi.ID
	for _, obj := range objects {
		relation := obj.Obj.(*topoapi.Object_Relation)
		nodeIDs = append(nodeIDs, relation.Relation.TgtEntityID)
	}
	nodeID := nodeIDs[0]

	cells, err := utils.GetCellIDsPerNode(nodeID)
	if err != nil {
		return err
	}
	cellObjectID := cells[0].CellObjectID

	numApps := sim.Arg("app-count").Int(1)
	numInstances := sim.Arg("replica-count").Int(1)
	numSubs := sim.Arg("sub-count").Int(1)

	subs := make([]simAppSub, numSubs)
	for i := 0; i < numSubs; i++ {
		subs[i] = simAppSub{
			name:         fmt.Sprintf("sub-%d", i+1),
			nodeID:       string(nodeID),
			cellObjectID: cellObjectID,
			reportPeriod: uint32((i + 1) * 5 * 1000),
			granularity:  500,
		}
	}

	s.apps = make([]*simApp, numApps)
	for i := 0; i < numApps; i++ {
		appID := fmt.Sprintf("sim-%s-%d-%d", sim.Name, sim.Process, i+1)
		instances := make([]*simAppInstance, numInstances)
		for j := 0; j < numInstances; j++ {
			instanceSubs := make([]*simAppSub, len(subs))
			for k := 0; k < len(subs); k++ {
				sub := subs[k]
				instanceSubs[k] = &sub
			}
			instances[j] = &simAppInstance{
				name:    fmt.Sprintf("%s-%d", appID, j),
				address: fmt.Sprintf("%s-%d.%s:%d", appID, j, appID, controlPort),
				subs:    instanceSubs,
			}
		}

		app := &simApp{
			name:      appID,
			instances: instances,
		}

		err := app.start()
		if err != nil {
			log.Error(err)
			return err
		}
		s.apps[i] = app
	}
	return nil
}

// ScheduleSimulator :: simulation
func (s *SimSuite) ScheduleSimulator(sim *simulation.Simulator) {
	sim.Schedule("subscribe", s.SimulateSubscribe, 1*time.Minute, 2)
	sim.Schedule("unsubscribe", s.SimulateUnsubscribe, 5*time.Minute, 1)
	sim.Schedule("crash", s.SimulateCrash, 10*time.Minute, 2)
}

func (s *SimSuite) getRandApp() *simApp {
	return s.apps[rand.Intn(len(s.apps))]
}

func (s *SimSuite) getRandInstance() *simAppInstance {
	app := s.getRandApp()
	instance := app.instances[rand.Intn(len(app.instances))]
	return instance
}

func (s *SimSuite) getRandSub(predicate func(sub *simAppSub) bool) (*simAppInstance, *simAppSub, bool) {
	var instances []*simAppInstance
	for _, app := range s.apps {
		for _, instance := range app.instances {
			for _, sub := range instance.subs {
				if predicate(sub) {
					instances = append(instances, instance)
					break
				}
			}
		}
	}

	if len(instances) == 0 {
		return nil, nil, false
	}

	instance := instances[rand.Intn(len(instances))]

	subs := make([]*simAppSub, 0, len(instance.subs))
	for _, sub := range instance.subs {
		if predicate(sub) {
			subs = append(subs, sub)
		}
	}

	if len(subs) == 0 {
		return nil, nil, false
	}

	sub := subs[rand.Intn(len(subs))]
	return instance, sub, true
}

func (s *SimSuite) SimulateSubscribe(sim *simulation.Simulator) error {
	instance, sub, ok := s.getRandSub(func(sub *simAppSub) bool {
		return !sub.open
	})
	if !ok {
		return nil
	}

	log.Infof("Starting '%s' subscription '%s' on '%s'", sub.nodeID, sub.name, instance.name)
	conn, err := grpc.Dial(instance.address, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return err
	}
	defer conn.Close()
	client := NewSimServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &StartSubscriptionRequest{
		SubscriptionId: sub.name,
		NodeId:         sub.nodeID,
		CellObjectId:   sub.cellObjectID,
		ReportPeriod:   sub.reportPeriod,
		Granularity:    sub.granularity,
	}
	_, err = client.StartSubscription(ctx, request)
	if err != nil {
		log.Error(err)
		return err
	}
	sub.open = true
	return nil
}

func (s *SimSuite) SimulateUnsubscribe(sim *simulation.Simulator) error {
	instance, sub, ok := s.getRandSub(func(sub *simAppSub) bool {
		return sub.open
	})
	if !ok {
		return nil
	}

	log.Infof("Stopping '%s' subscription '%s' on '%s'", sub.nodeID, sub.name, instance.name)
	conn, err := grpc.Dial(instance.address, grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return err
	}
	defer conn.Close()
	client := NewSimServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &StopSubscriptionRequest{
		SubscriptionId: sub.name,
		NodeId:         sub.nodeID,
	}
	_, err = client.StopSubscription(ctx, request)
	if err != nil {
		log.Error(err)
		return err
	}
	sub.open = false
	return nil
}

func (s *SimSuite) SimulateCrash(sim *simulation.Simulator) error {
	instance := s.getRandInstance()
	log.Infof("Crashing pod '%s'", instance.name)
	client, err := kubernetes.New()
	if err != nil {
		log.Error(err)
		return err
	}

	err = client.Clientset().
		CoreV1().
		Pods(client.Namespace()).
		Delete(instance.name, &metav1.DeleteOptions{})
	if err != nil {
		log.Error(err)
		return err
	}

	for _, sub := range instance.subs {
		sub.open = false
	}
	return nil
}

type simApp struct {
	name      string
	instances []*simAppInstance
}

func (s *simApp) start() error {
	log.Infof("Starting app '%s'", s.name)
	client, err := kubernetes.New()
	if err != nil {
		log.Error(err)
		return err
	}

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

	_, err = client.Clientset().
		AppsV1().
		StatefulSets(client.Namespace()).
		Create(ss)
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

	_, err = client.Clientset().
		CoreV1().
		Services(client.Namespace()).
		Create(svc)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

type simAppInstance struct {
	name    string
	address string
	subs    []*simAppSub
}

type simAppSub struct {
	name         string
	nodeID       string
	cellObjectID string
	reportPeriod uint32
	granularity  uint32
	open         bool
}
