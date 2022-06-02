import os
import sys
import subprocess
import shlex
import platform

__version__ = "0.5.0"

def run_tinygo():
    dir_name = os.path.dirname(os.path.realpath(__file__))
    dir_name = os.path.join(dir_name, "tinygo")
    tinygo = os.path.join(dir_name, 'bin/tinygo')
    if len(sys.argv) <= 1:
        return subprocess.call(tinygo, stdout=sys.stdout, stderr=sys.stderr)
    elif sys.argv[1] == "build":
        cmd = [tinygo]
        args = 'build -gc=leaking -target eosio -wasm-abi=generic -scheduler=none -opt z -tags=math_big_pure_go'
        args = shlex.split(args)
        cmd.extend(args)
        cmd.extend(sys.argv[2:])
        return subprocess.call(cmd, stdout=sys.stdout, stderr=sys.stderr)
    else:
        cmd = sys.argv[:]
        cmd[0] = tinygo
        return subprocess.call(cmd, stdout=sys.stdout, stderr=sys.stderr)

def run_command(tool):
    dir_name = os.path.dirname(os.path.realpath(__file__))
    dir_name = os.path.join(dir_name, "tinygo")
    tinygo = os.path.join(dir_name, 'bin/tinygo')
    cmd = sys.argv[:]
    cmd[0] = tinygo
    if platform.system() == 'Windows':
        tool += '.exe'
    cmd.insert(1, tool)
    return subprocess.call(cmd, stdout=sys.stdout, stderr=sys.stderr)

def run_clang():
    return run_command('clang')

def run_wasm_ld():
    return run_command('wasm-ld')

def run_ld_lld():
    return run_command('ld.lld')

def run_dlltool():
    return run_command('dlltool')

def run_ranlib():
    return run_command('ranlib')

def run_lib():
    return run_command('lib')

def run_ar():
    return run_command('ar')

def run_eosio_cpp():
    dir_name = os.path.dirname(os.path.realpath(__file__))
    dir_name = os.path.join(dir_name, "eosio.cdt")
    eosio_cpp = os.path.join(dir_name, 'bin/eosio-cpp')
    cmd = sys.argv[:]
    cmd[0] = eosio_cpp
    return subprocess.call(cmd, stdout=sys.stdout, stderr=sys.stderr)
