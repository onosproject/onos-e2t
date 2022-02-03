// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package oid

import "strings"

// Oid oid type
type Oid int

const (
	dotDelimiter = "."
)

const (
	iso                    = "iso"
	identifiedOrganization = "identifiedOrganization"
	dod                    = "dod"
	internet               = "internet"
	private                = "private"
	enterprise             = "enterprise"
	oran                   = "oran"
	version1               = "v1"
	version2               = "v2"
	e2smKpm                = "oran-e2sm-kpm"
	e2smRcPre              = "oran-e2sm-rc-pre"
	e2smMho                = "oran-e2sm-mho"
	e2smRsm                = "oran-e2sm-rsm"
	e2smNi                 = "oran-e2sm-ni"
	e2                     = "e2"
	e2sm                   = "e2sm"
)

const (
	isoOid                    Oid = 1
	identifiedOrganizationOid Oid = 3
	dodOid                    Oid = 6
	internetOid               Oid = 1
	privateOid                Oid = 4
	enterpriseOid             Oid = 1
	oranOid                   Oid = 53148
	e2Oid                     Oid = 1
	e2smOid                   Oid = 2
	version1Oid               Oid = 1
	version2Oid               Oid = 2
	e2smNiOid                 Oid = 1
	e2smKpmOid                Oid = 2
	e2smRcPreOid              Oid = 100
	e2smMhoOid                Oid = 101
	e2smRsmOid                Oid = 102
)

// defaultNodes a map of default oid names to their values
var defaultNodes = map[string]Oid{
	iso:                    isoOid,
	identifiedOrganization: identifiedOrganizationOid,
	dod:                    dodOid,
	internet:               internetOid,
	private:                privateOid,
	enterprise:             enterpriseOid,
	oran:                   oranOid,
	version1:               version1Oid,
	version2:               version2Oid,
	e2smKpm:                e2smKpmOid,
	e2smRcPre:              e2smRcPreOid,
	e2smMho:                e2smMhoOid,
	e2smRsm:                e2smRsmOid,
	e2smNi:                 e2smNiOid,
	e2:                     e2Oid,
	e2sm:                   e2smOid,
}

func createDottedOid(oidList []string) string {
	return strings.Join(oidList, dotDelimiter)
}
