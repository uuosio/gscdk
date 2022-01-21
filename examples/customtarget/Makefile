all : test

build : main.go
	eosio-go build -target ./custom.json -o test.wasm

test : build
	run-ipyeos -m pytest -x -s test.py

