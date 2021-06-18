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
	"google.golang.org/grpc"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math/rand"
	"time"
)

type simApp struct {
	name      string
	running   bool
	instances []*simAppInstance
}

type simAppInstance struct {
	name string
	subs []*simAppSub
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
	return helm.Chart("sd-ran").
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
}

// SetupSimulator :: simulation
func (s *SimSuite) SetupSimulator(sim *simulation.Simulator) error {
	err := helm.
		Chart("ran-simulator").
		Release(sim.Name).
		Install(true)
	if err != nil {
		return err
	}

	objects, err := utils.GetControlRelationObjects()
	if err != nil {
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
		appID := fmt.Sprintf("sim-%s-%d", sim.Name, i+1)
		instances := make([]*simAppInstance, numInstances)
		for j := 0; j < numInstances; j++ {
			instanceSubs := make([]*simAppSub, len(subs))
			for k := 0; k < len(subs); k++ {
				sub := subs[k]
				instanceSubs[k] = &sub
			}
			instances[j] = &simAppInstance{
				name: fmt.Sprintf("%s-%d", appID, j),
				subs: instanceSubs,
			}
		}
		s.apps[i] = &simApp{
			name:      appID,
			instances: instances,
		}
	}
	return nil
}

// ScheduleSimulator :: simulation
func (s *SimSuite) ScheduleSimulator(sim *simulation.Simulator) {
	sim.Schedule("start-app", s.SimulateStartApp, 10*time.Minute, .5)
	sim.Schedule("stop-app", s.SimulateStopApp, 1*time.Hour, .3)
	sim.Schedule("start-sub", s.SimulateStartSub, 10*time.Minute, .5)
	sim.Schedule("stop-sub", s.SimulateStopSub, 30*time.Minute, .3)
	sim.Schedule("crash-instance", s.SimulateCrashInstance, 30*time.Minute, .8)
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

	client, err := kubernetes.New()
	if err != nil {
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
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": app.name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "app",
							Image: "onosproject/onos-e2t-sim-app:latest",
							Args: []string{
								app.name,
								"$(POD_NAME)",
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
					Name: "sim",
					Port: 5000,
				},
			},
		},
	}

	_, err = client.Clientset().
		CoreV1().
		Services(client.Namespace()).
		Create(svc)
	if err != nil {
		return err
	}
	app.running = false
	return nil
}

func (s *SimSuite) SimulateStopApp(sim *simulation.Simulator) error {
	app, ok := s.getRunningApp()
	if !ok {
		return nil
	}

	client, err := kubernetes.New()
	if err != nil {
		return err
	}

	propagate := metav1.DeletePropagationForeground
	err = client.Clientset().
		AppsV1().
		StatefulSets(client.Namespace()).
		Delete(app.name, &metav1.DeleteOptions{PropagationPolicy: &propagate})
	if err != nil {
		return err
	}

	err = client.Clientset().
		CoreV1().
		Services(client.Namespace()).
		Delete(app.name, &metav1.DeleteOptions{PropagationPolicy: &propagate})
	if err != nil {
		return err
	}
	app.running = true
	return nil
}

func (s *SimSuite) SimulateStartSub(sim *simulation.Simulator) error {
	instance, sub, ok := s.getClosedSub()
	if !ok {
		return nil
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:5000", instance.name), grpc.WithInsecure())
	if err != nil {
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

	conn, err := grpc.Dial(fmt.Sprintf("%s:5000", instance.name), grpc.WithInsecure())
	if err != nil {
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

	client, err := kubernetes.New()
	if err != nil {
		return err
	}

	return client.Clientset().
		CoreV1().
		Pods(client.Namespace()).
		Delete(instance.name, &metav1.DeleteOptions{})
}