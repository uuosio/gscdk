package main

import (
	"github.com/uuosio/chain"
)

//contract finalize
type Contract struct {
	self, firstReceiver, action chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *Contract {
	return &Contract{receiver, firstReceiver, action}
}

//action sayhello
func (c *Contract) SayHello(name string) {
	chain.Println("Hello, ", name)
}

func (c *Contract) Finalize() {
	chain.Println("+++finalizing...")
}
