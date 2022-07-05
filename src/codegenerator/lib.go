package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"unsafe"

	traceable_errors "github.com/go-errors/errors"
)

//#include <stdint.h>
//#include <stddef.h>
//#include <stdbool.h>
// static char* get_p(char **pp, int i)
// {
//	    return pp[i];
// }
//typedef char *(*fn_malloc)(uint64_t size);
//static fn_malloc g_malloc = NULL;
//static void set_malloc_fn(fn_malloc fn) {
//	g_malloc = fn;
//}
//static char *cmalloc(uint64_t size) {
//	return 	g_malloc(size);
//}
import "C"

func renderData(data interface{}) *C.char {
	ret := map[string]interface{}{"data": data}
	result, _ := json.Marshal(ret)
	return CString(string(result))
}

func renderError(err error) *C.char {
	if _err, ok := err.(*traceable_errors.Error); ok {
		errMsg := _err.ErrorStack()
		ret := map[string]interface{}{"error": errMsg}
		result, _ := json.Marshal(ret)
		return CString(string(result))
	} else {
		pc, fn, line, _ := runtime.Caller(1)
		errMsg := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		ret := map[string]interface{}{"error": errMsg}
		result, _ := json.Marshal(ret)
		return CString(string(result))
	}
}

//export init_
func init_(malloc C.fn_malloc) {
	C.set_malloc_fn(malloc)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func CString(s string) *C.char {
	p := C.cmalloc(C.uint64_t(len(s) + 1))
	pp := (*[1 << 30]byte)(unsafe.Pointer(p))
	copy(pp[:], s)
	pp[len(s)] = 0
	return (*C.char)(p)
}

//func GenerateCode(inFile string, outFile string, tags []string) error
//export generate_code
func generate_code(inFile *C.char, outFile *C.char, tags **C.char, tags_length C.size_t) *C.char {
	_tags := []string{}
	for i := 0; i < int(tags_length); i++ {
		p := C.get_p(tags, C.int(i))
		_tags = append(_tags, C.GoString(p))
	}

	err := GenerateCode(C.GoString(inFile), C.GoString(outFile), _tags)
	if err == nil {
		return renderData("ok")
	} else {
		return renderError(err)
	}
}
