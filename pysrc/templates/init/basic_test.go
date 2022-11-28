package main

import (
	"context"
	"os"
	"testing"

	"github.com/uuosio/chaintester"
)

var ctx = context.Background()

func initTest() *chaintester.ChainTester {
	tester := chaintester.NewChainTester()
	testCoverage := os.Getenv("TEST_COVERAGE")
	if testCoverage == "TRUE" || testCoverage == "true" || testCoverage == "1" {
		tester.SetNativeApply("hello", ContractApply)
	}
	return tester
}

func TestHello(t *testing.T) {
	permissions := `
	{
		"hello": "active"
	}
	`

	tester := initTest()
	defer tester.FreeChain()

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

	err := tester.DeployContract("hello", "{{name}}.wasm", "{{name}}.abi")
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

	_, err = tester.PushAction("hello", "inc", "", permissions)
	if err != nil {
		panic(err)
	}
	// panic(ret.ToString())
	tester.ProduceBlock()
}
