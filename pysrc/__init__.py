import os
import sys
import subprocess
import shlex
import shutil
import platform
import argparse
import string
from .wasm_checker import check_import_section

__version__ = "0.7.4"

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

gscdk_install_dir = os.path.dirname(os.path.realpath(__file__))

def print_err(msg):
    print(f'{FAIL}:{msg}{ENDC}')

def print_warning(msg):
    print(f'{WARNING}:{msg}{ENDC}')

def find_wasm_file():
    return os.path.basename(os.path.realpath(os.curdir)) + '.wasm'

def gen_code(output, tags, package_name="main"):
    dir_name = os.path.dirname(os.path.realpath(__file__))
    dir_name = os.path.join(dir_name, "tinygo")
    code_generator = os.path.join(dir_name, 'bin/codegenerator')
    cmd = [code_generator, '-o', output, '-tags', tags, '-p', package_name]
    ret_code = subprocess.call(cmd, stdout=sys.stdout, stderr=sys.stderr)
    if not ret_code == 0:
        sys.exit(ret_code)

def build(wasm, tags, optimize=True):
    dir_name = os.path.dirname(os.path.realpath(__file__))
    dir_name = os.path.join(dir_name, "tinygo")
    tinygo = os.path.join(dir_name, 'bin/tinygo')
    if len(sys.argv) <= 1:
        return subprocess.call(tinygo, stdout=sys.stdout, stderr=sys.stderr)
    elif sys.argv[1] == "build":
        cmd = [tinygo]
        args = 'build -gc=leaking -target eosio -gen-code=false -wasm-abi=generic -scheduler=none -opt z -tags=math_big_pure_go'
        args = shlex.split(args)
        cmd.extend(args)
        cmd.extend(sys.argv[2:])
        ret_code = subprocess.call(cmd, stdout=sys.stdout, stderr=sys.stderr)
        if not ret_code == 0:
            sys.exit(ret_code)

        try:
            check_import_section(f'{wasm}')
        except Exception as e:
            print_err(f'{e}')
            sys.exit(-1)

        if not optimize:
            return
        
        wasm_opt = os.path.join(gscdk_install_dir, "binaryen-version_109/bin/wasm-opt")

        if 'Windows' == platform.system():
            wasm_opt += ".exe"
        if os.path.exists(wasm_opt):
            cmd = f'{wasm} -Oz --strip-debug -o {wasm}2'
            cmd = shlex.split(cmd)
            cmd.insert(0, wasm_opt)
            print(cmd)
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

def init(project_name):
    for c in project_name:
        if not c in string.ascii_lowercase or c == '_':
            raise Exception('project name contain invalid character(s), only characters in [a-z_] supported')
    os.mkdir(project_name)
    src_dir = os.path.dirname(__file__)
    for file in ['main.go', 'structs.go', 'tables.go', 'utils.go', 'basic_test.go', 'test.py', 'test.sh', 'build.sh', 'test.sh']:
        src_file = os.path.join(src_dir, f'templates/init/{file}')
        with open(src_file, 'r') as f:
            content = f.read()
        content = content.replace('{{name}}', project_name)
        file_path = f'{project_name}/{file}'
        with open(file_path, 'w') as f:
            f.write(content)
        if file.endswith('.sh'):
            if not 'Windows' == platform.system():
                os.chmod(file_path, 0o755)
    os.chdir(project_name)
    cmd = shlex.split(f'go mod init {project_name}')
    subprocess.call(cmd)

    cmd = shlex.split('go mod tidy')
    subprocess.call(cmd)
    
    gen_code("generated.go", "")

def run_tinygo():
    parser = argparse.ArgumentParser()
    subparsers = parser.add_subparsers(dest='subparser')

    sub_parser = subparsers.add_parser('init')
    sub_parser.add_argument('project_name')

    sub_parser = subparsers.add_parser('gencode')
    sub_parser.add_argument('-o', '--output', default="generated.go", help='ouput go file')
    sub_parser.add_argument('-tags', '--tags', type=str, default="", help='specify build tags')
    sub_parser.add_argument('-p', '--package-name', default="main", help='package name of generated code')

    sub_parser = subparsers.add_parser('build')
    sub_parser.add_argument('-o', '--output', help='target wasm file')
    sub_parser.add_argument('target', metavar='N', type=str, nargs='?', help='target wasm name')
    sub_parser.add_argument('-d', '--debug', action='store_true', help='enable debug build')
    sub_parser.add_argument('-gen-code', '--gen-code', type=str, default="true", help='enable code generation')
    sub_parser.add_argument('-tags', '--tags', type=str, default="", help='specify build tags')

    result = parser.parse_args()
    if not result:
        parser.print_usage()
        return

    if result.subparser == "init":
        init(result.project_name)
    elif result.subparser == "gencode":
        gen_code(result.output, result.tags, result.package_name)
    elif result.subparser == "build":
        if result.output:
            wasm = result.output
        else:
            wasm = find_wasm_file()

        tags = result.tags
        if result.gen_code == "true":
            gen_code('generated.go', tags)

        if result.debug:
            if '-d' in sys.argv: sys.argv.remove('-d')
            if '--debug' in sys.argv: sys.argv.remove('--debug')
            build(wasm, tags, False)
        else:
            build(wasm, tags, True)
    else:
        parser.print_usage()

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
