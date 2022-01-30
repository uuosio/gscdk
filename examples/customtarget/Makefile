all : test.wasm

test.wasm : main.go
	eosio-go build -target ./custom.json -o test.wasm

test : test.wasm
	run-ipyeos -m pytest -x -s test.py

