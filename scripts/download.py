import os
import sys
import time
import subprocess

version = sys.argv[1]
files = [
    f'gscdk-{version}-py3-none-macosx_10_15_x86_64.whl',
    f'gscdk-{version}-py3-none-manylinux1_x86_64.whl',
    f'gscdk-{version}-py3-none-win_amd64.whl',
]

url = f'https://github.com/uuosio/uuosio.gscdk/releases/download/v{version}/'
for f in files:
    count = 60*60/10
    while True:
        print('Downloading {}'.format(f))
        subprocess.call(['wget', url + f])
        if os.path.exists(f):
            break
        time.sleep(10)
        count -= 1
        if count <= 0:
            break

