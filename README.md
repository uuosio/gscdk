Go Smart Contracts Development Kit

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



# Debugging

Install `ipyeos` first for debugging.

```bash
python3 -m pip install ipyeos
```

In order to update to a new version, use the following command:

```bash
python3 -m pip install -U ipyeos
```

Then run the debugging server:

```bash
eos-debugger
```

On Windows, you need to use a docker image to run a debugging server.

```bash
docker pull ghcr.io/uuosio/ipyeos:latest
```

Run the debugging server on the Windows platform:

```bash
docker run -it --rm -p 9090:9090 -p 9092:9092 -t ghcr.io/uuosio/ipyeos
```

![Debugging](https://github.com/uuosio/gscdk/blob/main/images/debugging.gif)

# Code Coverage Analysis

Use the following command to generate a code coverage report in html

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

![Code Coverage](https://github.com/uuosio/gscdk/blob/main/images/code-coverage.png)

