// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package config

import (
	"github.com/onosproject/onos-lib-go/pkg/atomix"
	configlib "github.com/onosproject/onos-lib-go/pkg/config"
)

var config *Config

// Config is the onos-e2sub configuration
type Config struct {
	// Atomix is the Atomix configuration
	Atomix atomix.Config `yaml:"atomix,omitempty"`
}

// GetConfig gets the onos-e2sub configuration
func GetConfig() (Config, error) {
	if config == nil {
		config = &Config{}
		if err := configlib.Load(config); err != nil {
			return Config{}, err
		}
	}
	return *config, nil
}

// GetConfigOrDie gets the onos-e2sub configuration or panics
func GetConfigOrDie() Config {
	config, err := GetConfig()
	if err != nil {
		panic(err)
	}
	return config
}
