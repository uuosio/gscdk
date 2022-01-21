package main

import (
	"github.com/uuosio/chain"
)

//table mytable
type MyData struct {
	primary uint64 //primary: t.primary
	name    string
}

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
	code := c.Receiver
	scope := code
	payer := c.Receiver
	mydb := NewMyDataDB(code, scope)
	primary := uint64(111)
	if it, data := mydb.Get(primary); it.IsOk() {
		if data.name != name {
			chain.Println("Welcome new friend:", name)
		} else {
			chain.Println("Welcome old friend", name)
		}
		data.name = name
		mydb.Update(it, data, payer)
	} else {
		chain.Println("Welcome new friend", name)
		data := &MyData{primary, name}
		mydb.Store(data, payer)
	}
}
