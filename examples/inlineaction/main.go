package main

import (
	"github.com/uuosio/chain"
)

//packer
type Transfer struct {
	From     chain.Name
	To       chain.Name
	Quantity chain.Asset
	Memo     string
}

func main() {
	self, _, action := chain.GetApplyArgs()
	if action == chain.NewName("test") {
		t := Transfer{
			From:     self,
			To:       chain.NewName("eosio"),
			Quantity: *chain.NewAsset(10000, chain.NewSymbol("EOS", 4)),
			Memo:     "hello",
		}
		chain.NewAction(
			&chain.PermissionLevel{self, chain.ActiveName},
			chain.NewName("eosio.token"),
			chain.NewName("transfer"),
			&t,
		).Send()
		chain.Println("action sent!")
	}
}
