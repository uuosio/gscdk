package main

import (
	"github.com/uuosio/chain"
	"github.com/uuosio/chain/logger"
)

//table mytable
type MyData struct {
	primary uint64 //primary: t.primary
	name    string
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
	code := c.Receiver
	scope := code
	payer := c.Receiver
	mydb := NewMyDataDB(code, scope)
	primary := uint64(111)
	it, data := mydb.Get(primary)
	if !it.IsOk() {
		logger.Println("Welcome new friend", name)
		data := &MyData{primary, name}
		mydb.Store(data, payer)
	} else {
		if data.name != name {
			logger.Println("Welcome new friend:", name)
		} else {
			logger.Println("Welcome old friend", name)
		}
		data.name = name
		mydb.Update(it, data, payer)
	}
}
