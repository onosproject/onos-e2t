// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ricapie2

import (
	"errors"

	ricapie2v1beta1 "github.com/onosproject/onos-e2t/api/ricapi/e2/v1beta1"
)

// appRequestHandler
func (s Server) appRequestHandler(appReq *ricapie2v1beta1.AppRequest) error {
	//header := appReq.Header
	req := appReq.Req
	switch req.(type) {
	case *ricapie2v1beta1.AppRequest_SubReq:
		log.Info("Handling subscription request", string(appReq.Payload))
		// TODO send it to the subscription manager
	case *ricapie2v1beta1.AppRequest_SubDelReq:
		log.Info("Handling subscription delete request")
		// TODO send it to the subscription manager
	default:
		return errors.New("handling of the request type %v is not supported")
	}

	return nil
}
