#tinygo build -x -gc=leaking -target eosio -wasm-abi=generic -scheduler=none -opt 0 -tags=math_big_pure_go -gen-code=true -strip=false -o test.wasm .
mkdir -p build
#tinygo clang --target=wasm32--wasi --sysroot=$SYSROOT -Oz -I$SYSROOT/include -I$SYSROOT/include/libc -I$SYSROOT/include/libcxx -I$SYSROOT/include/eosiolib/capi -I$SYSROOT/include/eosiolib/core -I$SYSROOT/include/eosiolib/contracts -g -I$(pwd) -MD -MV -MTdeps -Xclang -internal-isystem -Xclang $SYSROOT/include/libc  -c -std=c++17 -Wno-unknown-attributes -o build/test.o test.cpp || exit 1
#tinygo clang --target=wasm32--wasi --sysroot=$SYSROOT -Oz -I$(pwd) -MD -MV -MTdeps -c -std=c++17 -Wno-unknown-attributes -internal-isystem -Xclang $SYSROOT/include/libc -o build/hello.o hello.cpp || exit 1
#-Xclang -internal-isystem -Xclang $SYSROOT/include/libc 
eosio-go build -o build/test.wasm . || exit 1
