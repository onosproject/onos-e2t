// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package controller

// Activator is an interface for controlling the activation of a controller
// Once the Activator is Started, it may activate or deactivate processing of Watcher events on the
// node at any time by writing true or false to the activator channel respectively.
type Activator interface {
	// Start starts the activator
	Start(ch chan<- bool) error

	// Stop stops the activator
	Stop()
}

// UnconditionalActivator activates controllers on all nodes at all times
type UnconditionalActivator struct {
}

// Start starts the activator
func (a *UnconditionalActivator) Start(ch chan<- bool) error {
	go func() {
		ch <- true
	}()
	return nil
}

// Stop stops the activator
func (a *UnconditionalActivator) Stop() {

}

var _ Activator = &UnconditionalActivator{}
