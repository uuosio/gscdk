package main

import (
	"flag"
	"strings"
)

func main() {
	_tags := flag.String("tags", "", "a space-separated list of extra build tags")
	output := flag.String("o", "", "output generated go file name")
	flag.Parse()

	tags := strings.Fields(*_tags)
	tags = append(tags, "eosio")
	tags = append(tags, "tinygo.wasm")
	err := GenerateCode(".", *output, tags)
	if err != nil {
		panic(err)
	}
}
