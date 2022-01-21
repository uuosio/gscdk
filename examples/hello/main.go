package main

import (
	"github.com/uuosio/chain"
)

//contract test
type MyContract struct {
	Receiver      chain.Name
	FirstReceiver chain.Name
	Action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *MyContract {
	return &MyContract{receiver, firstReceiver, action}
}

//action test
func (c *MyContract) Test(name string) {
	chain.Println("Hello", name)
}
