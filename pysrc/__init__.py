import os
import sys
import subprocess
import shlex

def run_tinygo():
    dir_name = os.path.dirname(os.path.realpath(__file__))
    dir_name = os.path.join(dir_name, "tinygo")
    tinygo = os.path.join(dir_name, 'bin/tinygo')
    if len(sys.argv) <= 1:
        subprocess.call(tinygo)
    elif sys.argv[1] == "build":
        cmd = [tinygo]
        args = 'build -x -gc=leaking -target eosio -wasm-abi=generic -scheduler=none -opt z -tags=math_big_pure_go'
        args = shlex.split(args)
        cmd.extend(args)
        cmd.extend(sys.argv[2:])
        subprocess.call(cmd)
    else:
        cmd = sys.argv[:]
        cmd[0] = tinygo
        subprocess.call(cmd)