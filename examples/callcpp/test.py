import os
import sys
import json
import time

from ipyeos import log
from ipyeos.chaintester import ChainTester

test_dir = os.path.dirname(__file__)
logger = log.get_logger(__name__)

tester = ChainTester()


def test_example():
    with open('test.wasm', 'rb') as f:
        code = f.read()
    tester.deploy_contract('hello', code, '', 0)
    r = tester.push_action('hello', 'test', '')
    logger.info(r['action_traces'][0]['console'])
    tester.produce_block()
