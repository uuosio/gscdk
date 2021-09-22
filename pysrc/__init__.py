import os
import sys
import subprocess
import shutil

def run_tinygo():
    dir_name = os.path.dirname(os.path.realpath(__file__))
    os.chdir(dir_name + "/tinygo/bin")

    if len(sys.argv) <= 1:
        cmd = ['tinygo']
        subprocess.call(cmd)
    elif sys.argv[1] == "build":
        cmd = 'tinygo build -x -gc=leaking -target eosio -wasm-abi=generic -scheduler=none -opt z -tags=math_big_pure_go'
        cmd = shutil.split(cmd)
        cmd.extend(sys.argv[2:])
        subprocess.call(cmd)
    else:
        cmd = ['tinygo']
        cmd.extend(sys.argv[1:])
        subprocess.call(cmd)
# if [ $1 == "build" ]; then
# tinygo build -x -gc=leaking -target eosio -wasm-abi=generic -scheduler=none -opt z -tags=math_big_pure_go ${@:2}
# else
# tinygo $@
# fi


