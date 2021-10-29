import os
import sys
import json
import time
from inspect import currentframe, getframeinfo

from uuoskit import wasmcompiler

test_dir = os.path.dirname(__file__)
sys.path.append(os.path.join(test_dir, '..'))

from uuosio import log
from uuosio.chaintester import ChainTester

logger = log.get_logger(__name__)

chain = ChainTester()

def test_hello():
    with open('hello.wasm', 'rb') as f:
        code = f.read()
    chain.deploy_contract('hello', code, b'', 0)

    r = chain.push_action('hello', 'test', b'')
    chain.produce_block()
