// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package subscription

import (
	"context"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
)

type Manager interface {
	Get(subID e2api.SubscriptionID) (Stream, error)
	Open(ctx context.Context, sub *e2api.Subscription) (Stream, error)
	Watch(ctx context.Context, ch <-chan Stream) error
}
