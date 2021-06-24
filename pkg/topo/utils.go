// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	"os"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

func getPodID() string {
	return os.Getenv("POD_ID")
}

func getE2CellRelationID(deviceID topoapi.ID, cellID topoapi.ID) (topoapi.ID, error) {
	cellRelationID := deviceID + "-" + cellID
	return cellRelationID, nil
}
