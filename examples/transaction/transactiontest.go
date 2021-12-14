package main

import (
	"github.com/uuosio/chain"
)

//contract test
type TransactionTest struct {
	Receiver      chain.Name
	FirstReceiver chain.Name
	Action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *TransactionTest {
	return &TransactionTest{receiver, firstReceiver, action}
}

//packer
type Transfer struct {
	From     chain.Name
	To       chain.Name
	Quantity chain.Asset
	Memo     string
}

//action sayhello
func (test *TransactionTest) SayHello() {
	payer := chain.NewName("helloworld11")

	t := Transfer{
		chain.NewName("helloworld11"),
		chain.NewName("eosio"),
		chain.Asset{10000, chain.NewSymbol("EOS", 4)},
		"hello,world",
	}

	a := chain.NewAction(
		chain.PermissionLevel{chain.NewName("helloworld11"), chain.ActiveName},
		chain.NewName("eosio.token"),
		chain.NewName("transfer"),
		&t,
	)

	tx := chain.NewTransaction(1)
	tx.Actions = []*chain.Action{a}
	tx.Send(1, false, payer)
	chain.Println("transaction sent")
}
