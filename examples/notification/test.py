
import os
import sys
try:
	from pyeoskit import eosapi, wallet
except:
	print('pyeoskit not found, please install it with "pip install pyeoskit"')
	sys.exit(-1)
from pyeoskit.exceptions import ChainException

# modify your test account here
test_account1 = 'helloworld11'
test_account2 = 'helloworld12'
# modify your test account private key here
wallet.import_key('test', '5JRYimgLBrRLCBAcjHUWCYRv3asNedTYYzVgmiU4q2ZVxMBiJXL')
wallet.import_key('test', '5JHRxntHapUryUetZgWdd3cg6BrpZLMJdqhhXnMaZiiT4qdJPhv')
# modify test node here
eosapi.set_node('https://testnode.uuos.network:8443')

with open('test.wasm', 'rb') as f:
    code = f.read()
abi = ''

try:
    eosapi.deploy_contract(test_account1, code, abi, vm_type=0)
except ChainException as e:
    if not e.json['error']['details'][0]['message'] == 'contract is already running this version of code':
        raise e

r = eosapi.transfer(test_account2, test_account1, 0.1, 'hello,world')
print(r['processed']['action_traces'][0]['inline_traces'][1]['console'])
balance = eosapi.get_balance(test_account1)
