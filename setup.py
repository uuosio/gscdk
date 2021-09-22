
import os
import setuptools
# from skbuild import setup
from distutils.core import setup

from distutils.sysconfig import get_python_lib
import glob

release_files = []
for root, dirs, files in os.walk("pysrc/tinygo"):
    for f in files:
        release_files.append(os.path.join(root.replace('pysrc/', ''), f))    

setup(
    name="gscdk",
    version="0.1.0",
    description="Go Smart Contract Development Kit",
    author='The UUOSIO Team',
    license="MIT",
    packages=['gscdk'],
    package_dir={'gscdk': 'pysrc'},
    package_data={
#        "": ["*"],
        'gscdk': release_files,
    },
    setup_requires=['wheel']
    # scripts=['compiler/build/release/tinygo/bin/eosio-go'],
    # install_requires=[
    # ],
    # include_package_data=True
)