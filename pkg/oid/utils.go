// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package oid

import (
	"strings"

	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

// ModelIDToOid converts service model ID to OID
func ModelIDToOid(r Registry, modelID string) (modelregistry.ModelOid, error) {
	log.Debugf("Converting service model ID %s to OID", modelID)
	id := strings.Split(modelID, dotDelimiter)
	if len(id) != 2 {
		return "", errors.NewInvalid("Invalid service model ID format", modelID)
	}

	smName := getOid(r, strings.ToLower(id[0]))
	version := getOid(r, strings.ToLower(id[1]))

	oidPrefix := createDottedOid([]string{getOid(r, iso),
		getOid(r, identifiedOrganization),
		getOid(r, dod),
		getOid(r, internet),
		getOid(r, private),
		getOid(r, enterprise),
		getOid(r, oran),
		getOid(r, e2)})

	modelOid := createDottedOid([]string{oidPrefix, version, getOid(r, e2sm), smName})
	return modelregistry.ModelOid(modelOid), nil

}
