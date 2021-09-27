<<<<<<< HEAD
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
=======
// SPDX-FileCopyrightText: ${year}-present Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0
>>>>>>> Implement E2 configuration controller and seperate management connection

package utils

import (
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/uri"
)

// GetE2TID gets E2T URI
func GetE2TID() topoapi.ID {
	return topoapi.ID(uri.NewURI(
		uri.WithScheme("e2"),
		uri.WithOpaque(env.GetPodID())).String())
}

func getFilter(kind string) *topoapi.Filters {
	controlRelationFilter := &topoapi.Filters{
		KindFilter: &topoapi.Filter{
			Filter: &topoapi.Filter_Equal_{
				Equal_: &topoapi.EqualFilter{
					Value: kind,
				},
			},
		},
	}
	return controlRelationFilter

}
func GetE2NodeFilter() *topoapi.Filters {
	return getFilter(topoapi.E2NODE)
}

func GetControlRelationFilter() *topoapi.Filters {
	return getFilter(topoapi.CONTROLS)
}

func GetE2TFilter() *topoapi.Filters {
	return getFilter(topoapi.E2T)
}
