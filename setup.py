
import os
import shutil
import setuptools
# from skbuild import setup
from distutils.core import setup

from distutils.sysconfig import get_python_lib
import glob

# if os.path.exists('pysrc/tinygo'):
#     shutil.rmtree('pysrc/tinygo')
# shutil.copytree('tinygo/build/release/tinygo', 'pysrc/tinygo')

release_files = []
for root, dirs, files in os.walk("pysrc/tinygo"):
    for f in files:
        release_files.append(os.path.join(root.replace('pysrc/', ''), f))    

# print(release_files)

setup(
    name="gscdk",
    version="0.5.0",
    description="Go Smart Contract Development Kit",
    author='The UUOSIO Team',
    license="BSD-3-Clause",
    url="https://github.com/uuosio/uuosio.gscdk",
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
