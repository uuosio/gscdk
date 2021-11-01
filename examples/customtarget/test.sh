eosio-go build -target ./custom.json -gen-code=true -o hello.wasm . || exit 1
run-uuos -m pytest -s -x test.py -k test_hello
