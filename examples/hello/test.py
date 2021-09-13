import os
from uuoskit import uuosapi, wallet, config
from uuoskit import log
logger = log.get_logger(__name__)

python_contract = config.python_contract
test_account1 = 'helloworld11'
test_account2 = 'helloworld12'

if os.path.exists('test.wallet'):
    os.remove('test.wallet')
psw = wallet.create('test')

# active key of helloworld11
wallet.import_key('test', '5JRYimgLBrRLCBAcjHUWCYRv3asNedTYYzVgmiU4q2ZVxMBiJXL', False)
wallet.import_key('test', '5Jbb4wuwz8MAzTB9FJNmrVYGXo4ABb7wqPVoWGcZ6x8V2FwNeDo', False)
# active key of helloworld12
wallet.import_key('test', '5JHRxntHapUryUetZgWdd3cg6BrpZLMJdqhhXnMaZiiT4qdJPhv', False)

uuosapi.set_node('https://testnode.uuos.network:8443')

with open('hello.go', 'r') as f:
    code = f.read()
code, abi = uuosapi.compile('hello', code, src_type=2)

try:
    uuosapi.deploy_contract(test_account1, code, abi, vm_type=0)
except Exception as e:
    print(e.json['error']['what'])

r = uuosapi.push_action(test_account1, 'sayhello', {'name': 'alice'}, {'hello':'active'})
print(r['processed']['action_traces'][0]['console'])
