#include <eosio/eosio.hpp>
extern "C" void say_hello() {
    eosio::print("++++++++hello,world\n");
}
