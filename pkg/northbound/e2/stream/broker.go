// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package stream

import "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/subscription"

type Broker interface {
	Transactions() TransactionManager
	Channels() ChannelManager
}

// NewBroker creates a new northbound stream broker
func NewBroker(subs subscription.Manager) (Broker, error) {
	transactions, err := newTransactionManager(subs)
	if err != nil {
		return nil, err
	}
	return &broker{
		transactions: transactions,
		channels:     newChannelManager(transactions),
	}, nil
}

type broker struct {
	channels     ChannelManager
	transactions TransactionManager
}

func (b *broker) Transactions() TransactionManager {
	return b.transactions
}

func (b *broker) Channels() ChannelManager {
	return b.channels
}
