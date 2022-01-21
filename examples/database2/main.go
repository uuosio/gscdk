package main

import (
	"github.com/uuosio/chain"
)

//table mydata2 singleton
type MySingleton struct {
	a1 uint64
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
func (t *MyContract) Test() {
	code := t.Receiver
	payer := t.Receiver
	db := NewMySingletonDB(code, code)

	data := db.Get()
	if data != nil {
		data.a1 += 1
		db.Set(data, payer)
	} else {
		data = &MySingleton{}
		db.Set(data, payer)
	}
	chain.Println("++data.a1:", data.a1)
}
