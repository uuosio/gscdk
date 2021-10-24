EOSIO_CDT_BUILD_DIR=eosio.cdt/build
EOSIO_CDT_RELEASE_DIR=pysrc/eosio.cdt

TINYGO_BUILD_DIR=tinygo/build
TINYGO_RELEASE_DIR=pysrc/tinygo

rm -r ${TINYGO_RELEASE_DIR}
rm -r ${TINYGO_BUILD_DIR}

mkdir -p ${EOSIO_CDT_RELEASE_DIR}/bin
mkdir -p ${EOSIO_CDT_RELEASE_DIR}/scripts
mkdir -p ${EOSIO_CDT_RELEASE_DIR}/include
mkdir -p ${EOSIO_CDT_RELEASE_DIR}/lib
mkdir -p ${EOSIO_CDT_RELEASE_DIR}/licenses

cp -R ${EOSIO_CDT_BUILD_DIR}/bin/* ${EOSIO_CDT_RELEASE_DIR}/bin || exit 1
cp -R ${EOSIO_CDT_BUILD_DIR}/licenses/* ${EOSIO_CDT_RELEASE_DIR}/licenses || exit 1
# install scripts
cp -R ${EOSIO_CDT_BUILD_DIR}/scripts/* ${EOSIO_CDT_RELEASE_DIR}/scripts  || exit 1
# install misc.
cp ${EOSIO_CDT_BUILD_DIR}/eosio.imports ${EOSIO_CDT_RELEASE_DIR} || exit 1

# install wasm includes
cp -R ${EOSIO_CDT_BUILD_DIR}/include/* ${EOSIO_CDT_RELEASE_DIR}/include || exit 1
# install wasm libs
cp ${EOSIO_CDT_BUILD_DIR}/lib/*.a ${EOSIO_CDT_RELEASE_DIR}/lib || exit 1

mkdir -p ${EOSIO_CDT_BUILD_DIR}/lib/cmake
cp -r ${EOSIO_CDT_BUILD_DIR}/lib/cmake/* ${EOSIO_CDT_RELEASE_DIR}/lib/cmake || exit 1

mkdir -p ${TINYGO_RELEASE_DIR}

mkdir -p ${TINYGO_BUILD_DIR}/release/tinygo/lib/eosio-libc/sysroot
mkdir -p ${TINYGO_BUILD_DIR}/release/tinygo/lib/eosio-libc/sysroot/include
mkdir -p ${TINYGO_BUILD_DIR}/release/tinygo/lib/eosio-libc/sysroot/lib

cp -r ${EOSIO_CDT_BUILD_DIR}/include/libc/* ${TINYGO_BUILD_DIR}/release/tinygo/lib/eosio-libc/sysroot/include
cp ${EOSIO_CDT_BUILD_DIR}/lib/*.a ${TINYGO_BUILD_DIR}/release/tinygo/lib/eosio-libc/sysroot/lib

cp -r ${TINYGO_BUILD_DIR}/release/tinygo/* ${TINYGO_RELEASE_DIR}
