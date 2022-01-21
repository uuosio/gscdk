package main

import (
	"github.com/uuosio/chain"
)

//contract test
type Contract struct {
	self, firstReceiver, action chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *Contract {
	return &Contract{receiver, firstReceiver, action}
}

//action test
func (c *Contract) Test() {
	db := NewCounterDB(c.self, c.self)
	item := db.Get()
	if item == nil {
		item = &Counter{count: 1}
	}
	item.count += 1
	db.Set(item, c.self)
	chain.Println("+++count:", item.count)
}
