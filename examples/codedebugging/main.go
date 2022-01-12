package main

import (
	"github.com/uuosio/chain"
)

//contract test
type Contract struct {
	self          chain.Name
	firstReceiver chain.Name
	action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *Contract {
	c := &Contract{receiver, firstReceiver, action}
	return c
}

//action sayhello
func (c *Contract) SayHello() {
	chain.Println("Hello, world!")
}
