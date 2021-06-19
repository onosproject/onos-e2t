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

type simApp struct {
	name      string
	running   bool
	instances []*simAppInstance
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
		s.apps[i] = &simApp{
			name:      appID,
			instances: instances,
		}
	}

	for _, app := range s.apps {
		err := s.startApp(sim, app)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

// ScheduleSimulator :: simulation
func (s *SimSuite) ScheduleSimulator(sim *simulation.Simulator) {
	sim.Schedule("start-app", s.SimulateStartApp, 5*time.Minute, 1)
	sim.Schedule("stop-app", s.SimulateStopApp, 30*time.Minute, 3)
	sim.Schedule("start-sub", s.SimulateStartSub, 1*time.Minute, 1)
	sim.Schedule("stop-sub", s.SimulateStopSub, 5*time.Minute, 1)
	sim.Schedule("crash-instance", s.SimulateCrashInstance, 10*time.Minute, 2)
}

func (s *SimSuite) getStoppedApp() (*simApp, bool) {
	stoppedApps := make([]*simApp, 0, len(s.apps))
	for _, app := range s.apps {
		if !app.running {
			stoppedApps = append(stoppedApps, app)
		}
	}
	if len(stoppedApps) == 0 {
		return nil, false
	}
	app := stoppedApps[rand.Intn(len(stoppedApps))]
	return app, true
}

func (s *SimSuite) getRunningApp() (*simApp, bool) {
	runningApps := make([]*simApp, 0, len(s.apps))
	for _, app := range s.apps {
		if app.running {
			runningApps = append(runningApps, app)
		}
	}
	if len(runningApps) == 0 {
		return nil, false
	}
	app := runningApps[rand.Intn(len(runningApps))]
	return app, true
}

func (s *SimSuite) getRunningInstance() (*simAppInstance, bool) {
	app, ok := s.getRunningApp()
	if !ok {
		return nil, false
	}
	instance := app.instances[rand.Intn(len(app.instances))]
	return instance, true
}

func (s *SimSuite) getClosedSub() (*simAppInstance, *simAppSub, bool) {
	instance, ok := s.getRunningInstance()
	if !ok {
		return nil, nil, false
	}

	closedSubs := make([]*simAppSub, 0, len(instance.subs))
	for _, sub := range instance.subs {
		if !sub.open {
			closedSubs = append(closedSubs, sub)
		}
	}
	if len(closedSubs) == 0 {
		return nil, nil, false
	}
	sub := closedSubs[rand.Intn(len(closedSubs))]
	return instance, sub, true
}

func (s *SimSuite) getOpenSub() (*simAppInstance, *simAppSub, bool) {
	instance, ok := s.getRunningInstance()
	if !ok {
		return nil, nil, false
	}

	openSubs := make([]*simAppSub, 0, len(instance.subs))
	for _, sub := range instance.subs {
		if sub.open {
			openSubs = append(openSubs, sub)
		}
	}
	if len(openSubs) == 0 {
		return nil, nil, false
	}
	sub := openSubs[rand.Intn(len(openSubs))]
	return instance, sub, true
}

func (s *SimSuite) SimulateStartApp(sim *simulation.Simulator) error {
	app, ok := s.getStoppedApp()
	if !ok {
		return nil
	}
	return s.startApp(sim, app)
}

func (s *SimSuite) startApp(sim *simulation.Simulator, app *simApp) error {
	log.Infof("Starting app '%s'", app.name)
	client, err := kubernetes.New()
	if err != nil {
		log.Error(err)
		return err
	}

	ss := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: app.name,
			Labels: map[string]string{
				"app": app.name,
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: app.name,
			Replicas:    pointer.Int32Ptr(int32(sim.Arg("replica-count").Int(1))),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": app.name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": app.name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "sim-app",
							Image:           "onosproject/onos-e2t-sim-app:latest",
							ImagePullPolicy: corev1.PullIfNotPresent,
							Args: []string{
								app.name,
								"$(POD_NAME)",
							},
							Env: []corev1.EnvVar{
								{
									Name: "NODE_ID",
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
			Name: app.name,
			Labels: map[string]string{
				"app": app.name,
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": app.name,
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
	app.running = true
	return nil
}

func (s *SimSuite) SimulateStopApp(sim *simulation.Simulator) error {
	app, ok := s.getRunningApp()
	if !ok {
		return nil
	}
	return s.stopApp(sim, app)
}

func (s *SimSuite) stopApp(sim *simulation.Simulator, app *simApp) error {
	log.Infof("Stopping app '%s'", app.name)
	client, err := kubernetes.New()
	if err != nil {
		log.Error(err)
		return err
	}

	propagate := metav1.DeletePropagationForeground
	err = client.Clientset().
		AppsV1().
		StatefulSets(client.Namespace()).
		Delete(app.name, &metav1.DeleteOptions{PropagationPolicy: &propagate})
	if err != nil {
		log.Error(err)
		return err
	}

	err = client.Clientset().
		CoreV1().
		Services(client.Namespace()).
		Delete(app.name, &metav1.DeleteOptions{PropagationPolicy: &propagate})
	if err != nil {
		log.Error(err)
		return err
	}

	app.running = false
	for _, instance := range app.instances {
		for _, sub := range instance.subs {
			sub.open = false
		}
	}
	return nil
}

func (s *SimSuite) SimulateStartSub(sim *simulation.Simulator) error {
	instance, sub, ok := s.getClosedSub()
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

func (s *SimSuite) SimulateStopSub(sim *simulation.Simulator) error {
	instance, sub, ok := s.getOpenSub()
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

func (s *SimSuite) SimulateCrashInstance(sim *simulation.Simulator) error {
	instance, ok := s.getRunningInstance()
	if !ok {
		return nil
	}

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
