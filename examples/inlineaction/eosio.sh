tinygo build  -x -gc=leaking -o eosio.wasm -target eosio -wasm-abi=generic -scheduler=none  -opt z -tags=math_big_pure_go .
if [ $? -ne 0 ]; then
    echo "build failed"
    exit 1
fi
eosio-wasm2wast -o eosio.wast eosio.wasm
eosio-wast2wasm -o eosio.wasm eosio.wast
