import os
import sys
import subprocess
import shlex

def run_tinygo():
    dir_name = os.path.dirname(os.path.realpath(__file__))
    dir_name = os.path.join(dir_name, "tinygo")
    os.environ['GOROOT'] = dir_name
    if len(sys.argv) <= 1:
        cmd = [f'{dir_name}/bin/tinygo']
        subprocess.call(cmd)
    elif sys.argv[1] == "build":
        cmd = f'{dir_name}/bin/tinygo build -x -gc=leaking -target eosio -wasm-abi=generic -scheduler=none -opt z -tags=math_big_pure_go'
        cmd = shlex.split(cmd)
        cmd.extend(sys.argv[2:])
        subprocess.call(cmd)
    else:
        cmd = [f'{dir_name}/bin/tinygo']
        cmd.extend(sys.argv[1:])
        subprocess.call(cmd)
