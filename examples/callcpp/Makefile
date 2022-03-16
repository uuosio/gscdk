all : test.wasm

test.wasm : main.go hello.cpp
	eosio-go build -o test.wasm

test : test.wasm
	run-ipyeos -m pytest -x -s test.py

