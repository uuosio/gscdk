#https://github.com/WebAssembly/design/blob/master/BinaryEncoding.md
section_types = (
    'custom',    #0
    'type',     #1
    'import',   #2
    'function', #3
    'table',    #4
    'memory',   #5
    'global',   #6
    'export',   #7
    'start',    #8
    'element',  #9
    'code',     #10
    'data'      #11
)

custom_section      = 0
type_section        = 1
import_section      = 2
function_section    = 3
table_section       = 4
memory_section      = 5
global_section      = 6
export_section      = 7
start_section       = 8
element_section     = 9
code_section        = 10
data_section        = 11

int_type = (
    127, #i32
    126, #i64
)

# external_kind
# A single-byte unsigned integer indicating the kind of definition being imported or defined:

# 0 indicating a Function import or definition
# 1 indicating a Table import or definition
# 2 indicating a Memory import or definition
# 3 indicating a Global import or definition

external_kinds = (
    'Function',
    'Table',
    'Memory',
    'Global'
)

allowed_functions = [
        "__ashrti3",
        "__lshlti3",
        "__lshrti3",
        "__ashlti3",
        "__divti3",
        "__udivti3",
        "__modti3",
        "__umodti3",
        "__multi3",
        "__addtf3",
        "__subtf3",
        "__multf3",
        "__divtf3",
        "__eqtf2",
        "__netf2",
        "__getf2",
        "__gttf2",
        "__lttf2",
        "__letf2",
        "__cmptf2",
        "__unordtf2",
        "__negtf2",
        "__floatsitf",
        "__floatunsitf",
        "__floatditf",
        "__floatunditf",
        "__floattidf",
        "__floatuntidf",
        "__floatsidf",
        "__extendsftf2",
        "__extenddftf2",
        "__fixtfti",
        "__fixtfdi",
        "__fixtfsi",
        "__fixunstfti",
        "__fixunstfdi",
        "__fixunstfsi",
        "__fixsfti",
        "__fixdfti",
        "__fixunssfti",
        "__fixunsdfti",
        "__trunctfdf2",
        "__trunctfsf2",
        "is_feature_active",
        "activate_feature",
        "get_resource_limits",
        "set_resource_limits",
        "set_proposed_producers",
        "get_blockchain_parameters_packed",
        "set_blockchain_parameters_packed",
        "is_privileged",
        "set_privileged",
        "get_active_producers",
        "db_idx64_store",
        "db_idx64_remove",
        "db_idx64_update",
        "db_idx64_find_primary",
        "db_idx64_find_secondary",
        "db_idx64_lowerbound",
        "db_idx64_upperbound",
        "db_idx64_end",
        "db_idx64_next",
        "db_idx64_previous",
        "db_idx128_store",
        "db_idx128_remove",
        "db_idx128_update",
        "db_idx128_find_primary",
        "db_idx128_find_secondary",
        "db_idx128_lowerbound",
        "db_idx128_upperbound",
        "db_idx128_end",
        "db_idx128_next",
        "db_idx128_previous",
        "db_idx256_store",
        "db_idx256_remove",
        "db_idx256_update",
        "db_idx256_find_primary",
        "db_idx256_find_secondary",
        "db_idx256_lowerbound",
        "db_idx256_upperbound",
        "db_idx256_end",
        "db_idx256_next",
        "db_idx256_previous",
        "db_idx_double_store",
        "db_idx_double_remove",
        "db_idx_double_update",
        "db_idx_double_find_primary",
        "db_idx_double_find_secondary",
        "db_idx_double_lowerbound",
        "db_idx_double_upperbound",
        "db_idx_double_end",
        "db_idx_double_next",
        "db_idx_double_previous",
        "db_idx_long_double_store",
        "db_idx_long_double_remove",
        "db_idx_long_double_update",
        "db_idx_long_double_find_primary",
        "db_idx_long_double_find_secondary",
        "db_idx_long_double_lowerbound",
        "db_idx_long_double_upperbound",
        "db_idx_long_double_end",
        "db_idx_long_double_next",
        "db_idx_long_double_previous",
        "db_store_i64",
        "db_update_i64",
        "db_remove_i64",
        "db_get_i64",
        "db_next_i64",
        "db_previous_i64",
        "db_find_i64",
        "db_lowerbound_i64",
        "db_upperbound_i64",
        "db_end_i64",
        "assert_recover_key",
        "recover_key",
        "assert_sha256",
        "assert_sha1",
        "assert_sha512",
        "assert_ripemd160",
        "sha1",
        "sha256",
        "sha512",
        "ripemd160",
        "check_transaction_authorization",
        "check_permission_authorization",
        "get_permission_last_used",
        "get_account_creation_time",
        "current_time",
        "publication_time",
        "abort",
        "eosio_assert",
        "eosio_assert_message",
        "eosio_assert_code",
        "eosio_exit",
        "read_action_data",
        "action_data_size",
        "current_receiver",
        "require_recipient",
        "require_auth",
        "require_auth2",
        "has_auth",
        "is_account",
        "prints",
        "prints_l",
        "printi",
        "printui",
        "printi128",
        "printui128",
        "printsf",
        "printdf",
        "printqf",
        "printn",
        "printhex",
        "read_transaction",
        "transaction_size",
        "expiration",
        "tapos_block_prefix",
        "tapos_block_num",
        "get_action",
        "send_inline",
        "send_context_free_inline",
        "send_deferred",
        "cancel_deferred",
        "get_context_free_data",
        "memcpy",
        "memmove",
        "memcmp",
        "memset",
]

class WasmReader(object):
    def __init__(self, raw):
        self.raw = raw
        self.idx = 0

    def read_bytes(self,size):
        ret = self.raw[self.idx:self.idx+size]
        self.idx += size
        return ret

    def read_byte(self):
        ret = self.raw[self.idx]
        self.idx += 1
        return ret

    def spec_binary_byte(self):
        if len(self.raw)<=self.idx:
            raise Exception("malformed")
        ret = self.raw[self.idx]
        self.idx += 1
        return ret

    #unsigned
    def spec_binary_uN(self, N):
        n=self.spec_binary_byte()
        if n<2**7 and n<2**N:
            return n
        elif n>=2**7 and N>7:
            m=self.spec_binary_uN(N-7)
            return (2**7)*m+(n-2**7)
        else:
            raise Exception("malformed")

    def read_u7(self):
        return self.spec_binary_uN(7)

    def read_u32(self):
        return self.spec_binary_uN(32)

    def end(self):
        return self.idx == len(self.raw)
    
    def remains(self):
        return self.raw[self.idx:]

    def read_uint32(self):
        return self.read_bytes(4)

    def read_uint64(self):
        return self.read_bytes(8)

def check_import_section(wasm_file):
    with open(wasm_file, 'rb') as f:
        raw = f.read()

    r = WasmReader(raw)

    magic = r.read_bytes(4)
    version = r.read_bytes(4)

    not_allowed_functions = []
    while not r.end():
        id = r.read_byte()
        payload_len = r.read_u32()
        payload = r.read_bytes(payload_len)
        pr = WasmReader(payload)
        count = pr.read_u32()
        if id == import_section:
            for i in range(count):
                module_len = pr.read_u32()
                module_str = pr.read_bytes(module_len)
                module_str = module_str.decode()

                field_len = pr.read_u32()
                field_str = pr.read_bytes(field_len)
                field_str = field_str.decode()

                kind = pr.read_u7()
                assert kind == 0
                type_index = pr.read_u32()
                if module_str != 'env' or not field_str in allowed_functions:
                    not_allowed_functions.append(f'{module_str}.{field_str}')
    if not_allowed_functions:
        raise Exception(f'imported function(s) not allowed: {not_allowed_functions}') 
