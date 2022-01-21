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
	_, _, action := chain.GetApplyArgs()
	if action == chain.NewName("sayhello") {
		t := Transfer{
			From:     chain.NewName("helloworld11"),
			To:       chain.NewName("eosio"),
			Quantity: *chain.NewAsset(10000, chain.NewSymbol("EOS", 4)),
			Memo:     "hello",
		}
		chain.NewAction(
			&chain.PermissionLevel{chain.NewName("helloworld11"), chain.ActiveName},
			chain.NewName("eosio.token"),
			chain.NewName("transfer"),
			&t,
		).Send()
	}
}
