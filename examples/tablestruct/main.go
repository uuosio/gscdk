package main

import "github.com/uuosio/chain"

//table mytable
type MyTable struct {
	id   uint64 //primary: t.id
	name string
}

//table mytable2 singleton
type MyTable2 struct {
	id   uint64
	name string
}

//ignore attribute denote this table will not include in ABI
//table mytable3 ignore
type MyTable3 struct {
	id   uint64 //primary: t.id
	name string
}

//singleton attribute can conbine with ignore attribute
//table mytable3 ignore singleton
type MyTable4 struct {
	id   uint64
	name string
}

//contract test
type MyContract struct {
	self          chain.Name
	firstReceiver chain.Name
	action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *MyContract {
	return &MyContract{
		self:          receiver,
		firstReceiver: firstReceiver,
		action:        action,
	}
}

//action sayhello
func (c *MyContract) SayHello(name string) {
	chain.Println("hello", name)
}
