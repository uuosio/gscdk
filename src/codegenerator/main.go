package main

import (
	"flag"
	"strings"
)

func main() {
	_tags := flag.String("tags", "", "a space-separated list of extra build tags")
	flag.Parse()

	tags := strings.Fields(*_tags)
	tags = append(tags, "eosio")
	tags = append(tags, "tinygo.wasm")
	err := GenerateCode(".", "generated.go", tags)
	if err != nil {
		panic(err)
	}
}
