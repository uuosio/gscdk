all : test.wasm

test.wasm : main.go hello.cpp
	go-contract build -o test.wasm .

test : test.wasm
	ipyeos -m pytest -x -s test.py

