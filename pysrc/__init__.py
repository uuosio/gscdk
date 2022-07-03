import os
import sys
import subprocess
import shlex
import shutil
import platform

__version__ = "0.6.1"

#https://stackabuse.com/how-to-print-colored-text-in-python/
#https://stackoverflow.com/questions/287871/how-do-i-print-colored-text-to-the-terminal
HEADER = '\033[95m'
OKBLUE = '\033[94m'
OKCYAN = '\033[96m'
OKGREEN = '\033[92m'
WARNING = '\033[1;30;43m'
FAIL = '\033[91m'
ENDC = '\033[0m'
BOLD = '\033[1m'
UNDERLINE = '\033[4m'

def print_err(msg):
    print(f'{FAIL}:{msg}{ENDC}')

def print_warning(msg):
    print(f'{WARNING}:{msg}{ENDC}')

def find_wasm_file():
    try:
        with open('go.mod') as f:
            return f.readlines()[0].strip().split(' ')[1] + '.wasm'
    except FileNotFoundError:
        return None

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
        ret_code = subprocess.call(cmd, stdout=sys.stdout, stderr=sys.stderr)
        if not ret_code == 0:
            sys.exit(ret_code)

        wasm = find_wasm_file()
        if shutil.which('wasm-opt'):
            cmd = f'wasm-opt {wasm} -Oz --strip-debug -o {wasm}2'
            cmd = shlex.split(cmd)
            ret_code = subprocess.call(cmd, stdout=sys.stdout, stderr=sys.stderr)
            if not ret_code == 0:
                sys.exit(ret_code)
            shutil.move(f"{wasm}2",  f"{wasm}")
        else:
            print_warning('''
wasm-opt not found! Make sure the binary is in your PATH environment.
We use this tool to optimize the size of your contract's Wasm binary.
wasm-opt is part of the binaryen package. You can find detailed
installation instructions on https://github.com/WebAssembly/binaryen#tools.
There are ready-to-install packages for many platforms:
* Debian/Ubuntu: apt-get install binaryen
* Homebrew: brew install binaryen
* Arch Linux: pacman -S binaryen
* Windows: binary releases at https://github.com/WebAssembly/binaryen/releases''')
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
