// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package oid

import (
	"errors"

	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
)

// ModelIDToOid converts service model ID to OID
func ModelIDToOid(r Registry, name string, version string) (types.OID, error) {
	log.Debugf("Converting service model ID %s:%s to the corresponding OID", name, version)
	if name == "" || version == "" {
		return "", errors.New("service model name and version must be specified")
	}
	oidPrefix := createDottedOid([]string{getOid(r, iso),
		getOid(r, identifiedOrganization),
		getOid(r, dod),
		getOid(r, internet),
		getOid(r, private),
		getOid(r, enterprise),
		getOid(r, oran),
		getOid(r, e2)})

	modelOid := createDottedOid([]string{oidPrefix, getOid(r, version), getOid(r, e2sm), getOid(r, name)})
	return types.OID(modelOid), nil

}
