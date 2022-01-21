package main

import (
	"github.com/uuosio/chain"
)

//contract test
type TransactionTest struct {
	self          chain.Name
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

//action test
func (test *TransactionTest) Test() {
	payer := test.self

	t := Transfer{
		test.self,
		chain.NewName("eosio"),
		chain.Asset{10000, chain.NewSymbol("EOS", 4)},
		"hello,world",
	}

	a := chain.NewAction(
		&chain.PermissionLevel{test.self, chain.ActiveName},
		chain.NewName("eosio.token"),
		chain.NewName("transfer"),
		&t,
	)

	tx := chain.NewTransaction(1)
	tx.Actions = []*chain.Action{a}
	tx.Send(chain.NewUint128(1, 0), false, payer)
	chain.Println("transaction sent")
}
