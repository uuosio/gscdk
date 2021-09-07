package main

import (
	"github.com/uuosio/chain"
)

//contract hello
type MyContract struct {
	Receiver      chain.Name
	FirstReceiver chain.Name
	Action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *MyContract {
	return &MyContract{receiver, firstReceiver, action}
}

//action sayhello
func (c *MyContract) SayHello(name string) {
	a := make([]string, 1)
	chain.Check(a[0] == "", a[0])
	chain.Println("Hello", name, a)
}
