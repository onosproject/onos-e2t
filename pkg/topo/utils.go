// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	"fmt"
	"os"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/uri"
)

func getPodID() string {
	return os.Getenv("POD_ID")
}

func GetCellID(e2NodeID topoapi.ID, cellGlobalID string) topoapi.ID {
	opaque := fmt.Sprintf("%s/%s", e2NodeID, cellGlobalID)
	// e2 scheme is already included in e2 node ID so not needed to be include in cell uri
	cellID := topoapi.ID(uri.NewURI(
		uri.WithOpaque(opaque)).String())
	return cellID
}

func GetE2TID() topoapi.ID {
	opaque := getPodID()
	return topoapi.ID(uri.NewURI(
		uri.WithScheme("e2"),
		uri.WithOpaque(opaque)).String())
}
