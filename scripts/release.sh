BUILD_DIR=eosio.cdt/build
EOSIO_CDT_RELEASE_DIR=pysrc/eosio.cdt
TINYGO_RELEASE_DIR=pysrc/tinygo

mkdir -p ${EOSIO_CDT_RELEASE_DIR}/bin
mkdir -p ${EOSIO_CDT_RELEASE_DIR}/scripts
mkdir -p ${EOSIO_CDT_RELEASE_DIR}/include
mkdir -p ${EOSIO_CDT_RELEASE_DIR}/lib
mkdir -p ${EOSIO_CDT_RELEASE_DIR}/licenses

cp -R ${BUILD_DIR}/bin/* ${EOSIO_CDT_RELEASE_DIR}/bin || exit 1
cp -R ${BUILD_DIR}/licenses/* ${EOSIO_CDT_RELEASE_DIR}/licenses || exit 1
# install scripts
cp -R ${BUILD_DIR}/scripts/* ${EOSIO_CDT_RELEASE_DIR}/scripts  || exit 1
# install misc.
cp ${BUILD_DIR}/eosio.imports ${EOSIO_CDT_RELEASE_DIR} || exit 1

# install wasm includes
cp -R ${BUILD_DIR}/include/* ${EOSIO_CDT_RELEASE_DIR}/include || exit 1
# install wasm libs
cp ${BUILD_DIR}/lib/*.a ${EOSIO_CDT_RELEASE_DIR}/lib || exit 1

mkdir -p ${BUILD_DIR}/lib/cmake
cp -r ${BUILD_DIR}/lib/cmake/* ${EOSIO_CDT_RELEASE_DIR}/lib/* || exit 1

mkdir -p ${TINYGO_RELEASE_DIR}

cp -r ./tinygo/build/release/tinygo/* ${TINYGO_RELEASE_DIR}
