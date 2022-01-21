package main

/*
#cgo CFLAGS: -I./include
#include "funcs.h"
*/
import "C"

func main() {
	C.say_hello()
}
