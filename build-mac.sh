pushd src/codegenerator
go build .
popd
cp src/codegenerator/codegenerator.exe ./pysrc/tinygo/bin
python3 setup.py sdist bdist_wheel --plat-name macosx_10_15_x86_64
