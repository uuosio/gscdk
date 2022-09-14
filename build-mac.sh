pushd src/codegenerator
go build .
popd
cp src/codegenerator/codegenerator ./pysrc/tinygo/bin || exit
python3 setup.py sdist bdist_wheel --plat-name macosx_10_15_x86_64
