package main

import (
	"github.com/uuosio/chain"
)

//contract {{name}}
type Contract struct {
	receiver      chain.Name
	firstReceiver chain.Name
	action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *Contract {
	return &Contract{
		receiver,
		firstReceiver,
		action,
	}
}

//action inc
func (c *Contract) Inc(name string) {
	db := NewCounterTable(c.receiver)
	it := db.Find(1)
	payer := c.receiver
	if it.IsOk() {
		value := db.GetByIterator(it)
		value.count += 1
		db.Update(it, value, payer)
		chain.Println("count: ", value.count)
	} else {
		value := &Counter{
			key:   1,
			count: 1,
		}
		db.Store(value, payer)
		chain.Println("count: ", value.count)
	}
}
