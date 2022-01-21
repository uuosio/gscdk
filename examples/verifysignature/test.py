import os
import sys
import json
import time
import hashlib

from ipyeos import log
from ipyeos.chaintester import ChainTester
from pyeoskit import eosapi

test_dir = os.path.dirname(__file__)
logger = log.get_logger(__name__)

tester = ChainTester()

def read_code_and_abi():
    files = os.listdir(test_dir)
    abi = None
    code = None
    for file in files:
        if file.endswith('.wasm'):
            file = os.path.join(test_dir, file)
            with open(file, 'rb') as f:
                code = f.read()
        elif file.endswith('.abi'):
            file = os.path.join(test_dir, file)
            with open(file, 'r') as f:
                abi = json.load(f)
    return code, abi

def test_example():
    code, abi = read_code_and_abi()
    tester.deploy_contract('hello', code, abi, 0)

    data = 'hello world'
    digest = hashlib.sha256(data.encode()).hexdigest()
    signature = eosapi.sign_digest(digest, '5KCbo1u7A36GAv7LRepcD1SSmkCVVNEGvpiY3cAn3X7ATDGnd7A')

    public_key = 'EOS4xM57qXUyvhGjBPf7YzwnLjQVCy8hLrqSJqBb8gHBZ2X8cyC6v'
    args = {
        'data': 'hello world',
        'public_key': public_key,
        'signature': signature,
    }
    r = tester.push_action('hello', 'verify', args)
    logger.info(r['action_traces'][0]['console'])
    tester.produce_block()
