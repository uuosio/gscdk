import os
import sys
import json
import time
import uuid
import hashlib

test_dir = os.path.dirname(__file__)
sys.path.append(os.path.join(test_dir, '..'))

from ipyeos import log
from ipyeos.chaintester import ChainTester
from pyeoskit import eosapi

logger = log.get_logger(__name__)
tester = ChainTester()

def test_event():
    with open('verifier.wasm', 'rb') as f:
        code = f.read()
    with open('verifier.abi', 'r') as f:
        abi = f.read()
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
    for i in range(5):
        r = tester.push_action('hello', 'verify', args)
        print(r['elapsed'])
        tester.produce_block()


