mkdir -p build
eosio-go build -o build/hello.wasm . || exit 1
