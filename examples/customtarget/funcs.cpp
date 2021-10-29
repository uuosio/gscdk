#include <eosio/eosio.hpp>

//compiler can find funcs.h, becuase custom.json already specify the include path
#include "funcs.h"

void say_hello() {
    eosio::print("hello ", NAME, "!\n");
}
