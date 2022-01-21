package main

import (
	"github.com/uuosio/chain"
)

//contract test
type MyContract struct {
	Receiver      chain.Name
	FirstReceiver chain.Name
	Action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *MyContract {
	return &MyContract{receiver, firstReceiver, action}
}

//action verify
func (c *MyContract) Verify(data string, public_key *chain.PublicKey, signature *chain.Signature) {
	digest := chain.Sha256([]byte(data))
	recovered_pub := chain.RecoverKey(digest, signature)
	chain.Assert(*recovered_pub == *public_key, "invalid signature")
	chain.Println("verify successful")
}
