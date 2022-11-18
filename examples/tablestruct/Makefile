all : test.wasm

test.wasm : main.go
	go-contract build -o test.wasm .

test : test.wasm
	ipyeos -m pytest -x -s test.py

