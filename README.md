Go Smart Contracts Development Kit

# What the Go Smart Contracts look like?

Here is an example

```go
package main

import (
	"github.com/uuosio/chain"
)

//table mytable
type MyData struct {
	primary uint64 //primary: t.primary
	name    string
}

//contract mycontract
type MyContract struct {
	Receiver      chain.Name
	FirstReceiver chain.Name
	Action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *MyContract {
	return &MyContract{receiver, firstReceiver, action}
}

//action sayhello
func (c *MyContract) SayHello(name string) {
	code := c.Receiver
	scope := code
	payer := c.Receiver
	mydb := NewMyDataDB(code, scope)
	primary := uint64(111)
	if it, data := mydb.Get(primary); it.IsOk() {
		if data.name != name {
			chain.Println("Welcome new friend:", name)
		} else {
			chain.Println("Welcome old friend", name)
		}
		data.name = name
		mydb.Update(it, data, payer)
	} else {
		chain.Println("Welcome new friend", name)
		data := &MyData{primary, name}
		mydb.Store(data, payer)
	}
}
```

# Installation

```bash
python3 -m pip install gscdk
```

For Windows platform:

```bash
python -m pip install gscdk
```

### Upgrade From an Old Version

```bash
python3 -m pip install --upgrade gscdk
```

For Windows platform:

```bash
python -m pip install --upgrade gscdk
```

### Installing gscdk in Docker

Building docker image for gscdk

```
docker build github.com/learnforpractice/gscdk-docker#main -t gscdk/test
```

Running

```
docker run -w /root/dev -it --rm -v "$(pwd)":/root/dev -t gscdk/test /bin/bash
```

# Building Go Smart Contracts Compiler

Follow the steps in [Building](./BUILDING.md)

That will build tinygo command in compiler/build directory that support for building Go Smart Contracts.

```bash
export PATH=$(pwd)/compiler/build:$PATH
```

# eosio-go

## Initializing a project with "init" subcommand

"init" command initialize a project with contract name

```
eosio-go init mycontract
cd mycontract
```

## Generating ABI and Extra Code for Smart Contracts

```
eosio-go gencode -o generated.go .
```

Code generation is also the default option for "build" command

## Building Go Smart Contracts Project

#### Compiling the Source Code

```bash
eosio-go build -o mycontract.wasm .
```

#### Disable Code Generation during Building

```bash
eosio-go build -gen-code=false -o mycontract .
```

#### Disable Striping WASM File after Building

```bash
eosio-go build -strip=false -o mycontract .
```
