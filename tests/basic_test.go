package main

import (
	"context"
	"github.com/uuosio/chain"
	"github.com/uuosio/chaintester"
	"testing"
)

var ctx = context.Background()

func OnApply(receiver, firstReceiver, action uint64) {
	n := chain.Name{action}
	println("+++++++++OnApply", n.String())
	contract_apply(receiver, firstReceiver, action)
}

func init() {
	chaintester.SetApplyFunc(OnApply)
}

func TestHello(t *testing.T) {
	// t.Errorf("++++++enable_debug: %v", os.Getenv("enable_debug"))
	permissions := `
	{
		"hello": "active"
	}
	`

	tester := chaintester.NewChainTester()
	defer tester.FreeChain()

	tester.EnableDebugContract("hello", true)

	updateAuthArgs := `{
		"account": "hello",
		"permission": "active",
		"parent": "owner",
		"auth": {
			"threshold": 1,
			"keys": [
				{
					"key": "EOS6AjF6hvF7GSuSd4sCgfPKq5uWaXvGM2aQtEUCwmEHygQaqxBSV",
					"weight": 1
				}
			],
			"accounts": [{"permission":{"actor": "hello", "permission": "eosio.code"}, "weight":1}],
			"waits": []
		}
	}`
	tester.PushAction("eosio", "updateauth", updateAuthArgs, permissions)

	err := tester.DeployContract("hello", "test.wasm", "test.abi")
	if err != nil {
		panic(err)
	}
	tester.ProduceBlock()

	_, err = tester.PushAction("hello", "inc", "", permissions)
	if err != nil {
		panic(err)
	}
	// panic(ret.ToString())
	tester.ProduceBlock()
}

func TestAssert(t *testing.T) {
	// t.Errorf("++++++enable_debug: %v", os.Getenv("enable_debug"))
	permissions := `
	{
		"hello": "active"
	}
	`

	tester := chaintester.NewChainTester()
	defer tester.FreeChain()

	tester.EnableDebugContract("hello", true)
	err := tester.DeployContract("hello", "test.wasm", "test.abi")
	if err != nil {
		panic(err)
	}
	tester.ProduceBlock()

	_, err = tester.PushAction("hello", "assert", "", permissions)
	if err == nil {
		panic("should return error")
	} else {
		_err, ok := err.(*chaintester.TransactionError)
		if !ok {
			panic("bad error")
		}
		t.Logf("++++++error: %v", _err)
	}
	// r.GetString("except")
	// panic(ret.ToString())
	tester.ProduceBlock()

	_, err = tester.PushAction("hello", "inc", "", permissions)
	if err != nil {
		panic(err)
	}
}
