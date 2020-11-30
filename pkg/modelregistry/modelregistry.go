// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package modelregistry

import (
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("modelregistry")

// ModelRegistry is the object for the saving information about device models
type ModelRegistry struct {
	ServiceModels map[string]ServiceModel
}

// Interface implemented by each Service Model Plugin
type ServiceModel interface {
	ServiceModelData() (string, string, string)
	IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error)
	IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error)
	IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error)
	IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error)
	RanFuncDescriptionASN1toProto([]byte) ([]byte, error)
	RanFuncDescriptionProtoToASN1([]byte) ([]byte, error)
}

func (registry *ModelRegistry) RegisterModelPlugin(moduleName string) (string, string, error) {
	log.Info("Loading module ", moduleName)

	return "", "", fmt.Errorf("not yet implemented")
}
