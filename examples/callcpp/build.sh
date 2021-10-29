mkdir -p build
eosio-go build -o build/test.wasm . || exit 1
