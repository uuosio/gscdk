package main

/*
#include <stdbool.h>
#include <stdint.h>

void say_hello();
bool find_db(uint64_t code, uint64_t scope, uint64_t key);
*/
import "C"
import "github.com/uuosio/chain"

func main() {
	receiver, code, _ := chain.GetApplyArgs()
	C.find_db(receiver.N, code.N, 1)
	C.say_hello()
}
