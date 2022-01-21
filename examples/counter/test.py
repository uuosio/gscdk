import os
import sys
import json
import time

from ipyeos import log
from ipyeos.chaintester import ChainTester

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
    r = tester.push_action('hello', 'test', '')
    logger.info(r['action_traces'][0]['console'])
    tester.produce_block()
