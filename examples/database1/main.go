package main

import (
	"github.com/uuosio/chain"
)

//table mytable
type MyData struct {
	primary uint64 //primary: t.primary
	n       uint64
}

//contract mycontract
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
	code := chain.NewName("hello")
	scope := code
	payer := code
	mydb := NewMyDataDB(code, scope)
	primary := uint64(1)
	it, data := mydb.Get(primary)
	if !it.IsOk() {
		data := &MyData{primary, 111}
		mydb.Store(data, payer)
	} else {
		data.n += 1
		mydb.Update(it, data, payer)
	}
}
