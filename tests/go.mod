module test

go 1.17

replace github.com/uuosio/chaintester => ../

require (
	github.com/uuosio/chain v0.2.1
	github.com/uuosio/chaintester v0.0.0-20220720033226-d528e3affc43
)

require (
	github.com/apache/thrift v0.16.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
)
