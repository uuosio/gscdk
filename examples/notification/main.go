package main

import (
	"github.com/uuosio/chain"
)

//contract test
type MyContract struct {
	Self          chain.Name
	FirstReceiver chain.Name
	Action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *MyContract {
	return &MyContract{receiver, firstReceiver, action}
}

//notify transfer
func (c *MyContract) Transfer(from, to chain.Name, quantity chain.Asset, memo string) {
	if to != c.Self {
		return
	}

	if c.FirstReceiver == chain.NewName("eosio.token") &&
		c.Action == chain.NewName("transfer") {
		if quantity.Symbol == chain.NewSymbol("EOS", 4) {
			chain.Println("Receive Eos, memo:", memo)
		}
	}
}
