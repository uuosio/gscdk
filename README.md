Go Smart Contracts Development Kit

# Debugger
![Debugging](https://github.com/uuosio/gscdk/blob/main/images/debugging.gif)


# Code Coverage
![Code Coverage](https://github.com/uuosio/gscdk/blob/main/images/code-coverage.png)

# What a Go Smart Contract looks like?

Here is an example

```go
package main

import (
	"github.com/uuosio/chain"
)

//table mytable
type MyData struct {
	primary uint64 //primary
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
	payer := c.Receiver
	mydb := NewMyDataDB(code)
	primary := uint64(111)
	if it, data := mydb.GetByKey(primary); it.IsOk() {
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

# Quick Start

[Quick Start](https://colab.research.google.com/github/uuosio/gscdk/blob/main/quickstart/quickstart.ipynb)

# Installation

```bash
python3 -m pip install gscdk
```

For the Windows platform:

```bash
python -m pip install gscdk
```

### Upgrade From an Old Version

```bash
python3 -m pip install --upgrade gscdk
```

For the Windows platform:

```bash
python -m pip install --upgrade gscdk
```

### Installing gscdk in Docker

Building docker image for gscdk

```
docker build https://github.com/uuosio/gscdk-docker#main -t gscdk/test
```

Running

```
docker run -w /root/dev -it --rm -v "$(pwd)":/root/dev -t gscdk/test /bin/bash
```

# Building Go Smart Contracts Compiler

Follow the steps in [Building](./BUILDING.md)

That will build the `tinygo` command in the compiler/build directory that supports building Go Smart Contracts.

```bash
export PATH=$(pwd)/compiler/build:$PATH
```

# go-contract

## Initializing a project with the "init" subcommand

The "init" command initializes a project with the contract name

```
go-contract init mycontract
cd mycontract
```

## Generating ABI and Extra Code for Smart Contracts

```
go-contract gencode
```

Code generation is also the default option for the "build" command

## Building Go Smart Contracts Project

#### Compiling the Source Code

```bash
go-contract build
```

#### Disable Code Generation during Building

```bash
go-contract build -gen-code=false .
```

#### Disable Code Optimization

Specifying `-d` or `--debug` option to disable wasm optimization.

```bash
go-contract build -d
```
