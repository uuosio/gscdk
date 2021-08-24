import os
from uuoskit import uuosapi, wallet, config

python_contract = config.python_contract
test_account1 = 'helloworld11'
test_account2 = 'helloworld12'

def init():
    if os.path.exists('test.wallet'):
        os.remove('test.wallet')
    psw = wallet.create('test')


    # active key of helloworld11
    wallet.import_key('test', '5JRYimgLBrRLCBAcjHUWCYRv3asNedTYYzVgmiU4q2ZVxMBiJXL')
    wallet.import_key('test', '5Jbb4wuwz8MAzTB9FJNmrVYGXo4ABb7wqPVoWGcZ6x8V2FwNeDo')
    # active key of helloworld12
    wallet.import_key('test', '5JHRxntHapUryUetZgWdd3cg6BrpZLMJdqhhXnMaZiiT4qdJPhv')

    config.python_contract = 'hello'
    config.main_token = 'EOS'
    config.main_token_contract = 'eosio.token'
    uuosapi.set_node('https://testnode.uuos.network:8443')

def run_test_code(code):
    code = uuosapi.mp_compile(python_contract, code)
    args = uuosapi.s2n(test_account1) + code
    uuosapi.push_action(python_contract, args, {test_account1:'active'})
    r = uuosapi.push_action(python_contract, 'exec', b'hello,world', {test_account1:'active'})
    print(r['processed']['action_traces'][0]['console'])
