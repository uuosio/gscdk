import os
import time
import subprocess

files = [
    'gscdk-0.3.0-py3-none-macosx_10_15_x86_64.whl',
    # 'gscdk-0.3.0-py3-none-manylinux1_x86_64.whl',
    'gscdk-0.3.0-py3-none-win_amd64.whl',
]

url = 'https://github.com/uuosio/uuosio.gscdk/releases/download/v0.3.0/'
for f in files:
    count = 60*60/10
    while not os.path.exists(f):
        print('Downloading {}'.format(f))
        subprocess.call(['wget', url + f])
        time.sleep(10)
        count -= 1
        if count <= 0:
            break

