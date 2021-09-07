package main

import (
	"github.com/uuosio/chain"
)

type TransactionTest struct {
	Receiver      chain.Name
	FirstReceiver chain.Name
	Action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *TransactionTest {
	return &TransactionTest{receiver, firstReceiver, action}
}

//action sayhello
func (test *TransactionTest) SayHello() {
	payer := chain.NewName("helloworld11")

	a := chain.Action{}
	a.Account = chain.NewName("eosio.token")
	a.Name = chain.NewName("transfer")
	a.AddPermission(chain.NewName("helloworld11"), chain.ActiveName)

	t := chain.Transfer{
		chain.NewName("helloworld11"),
		chain.NewName("eosio"),
		chain.Asset{10000, chain.NewSymbol("EOS", 4)},
		"hello,world",
	}
	a.Data = t.Pack()

	tx := chain.NewTransaction(1)
	tx.Actions = []chain.Action{a}
	tx.Send(1, false, payer)
	chain.Println("transaction sent")
}
