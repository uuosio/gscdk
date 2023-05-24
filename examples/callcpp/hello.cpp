#include <eosio/eosio.hpp>
#include <eosio/multi_index.hpp>

using namespace eosio;

struct record {
   uint64_t    primary;
   uint128_t   secondary;
   uint64_t    data;
   uint64_t primary_key() const { return primary; }
   uint128_t get_secondary() const { return secondary; }
   EOSLIB_SERIALIZE( record, (primary)(secondary)(data))
};

extern "C" void say_hello() {
    uint64_t *ptr = new uint64_t(0);
    eosio::print(uint64_t(ptr), "\n");
    eosio::print_f("++++++++hello,world!!! %\n", uint64_t(ptr));
    delete ptr;
}


extern "C" bool find_db(uint64_t code, uint64_t scope, uint64_t key) {
    multi_index<"mytable4"_n,
                record,
                indexed_by< "bysecondary"_n,
                const_mem_fun<record, uint128_t, &record::get_secondary> > > mytable(name{code}, scope);
    auto it = mytable.find(key);
    return it != mytable.end();
}

