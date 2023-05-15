# Go Smart Contracts Development Kit (GSCDK)

[![PyPi Version](https://img.shields.io/pypi/v/gscdk.svg)](https://pypi.org/project/gscdk)
[![PyPi Downloads](https://img.shields.io/pypi/dm/gscdk.svg)](https://pypi.org/project/gscdk)

## Overview

The Go Smart Contracts Development Kit (GSCDK) provides a comprehensive toolkit for creating, building, and debugging Go-based smart contracts. 

## Example of a Go Smart Contract

Here is an example of what a Go Smart Contract looks like using GSCDK:

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


## Quick Start

Jump right into building your first smart contract with our [Quick Start Guide](https://colab.research.google.com/github/uuosio/gscdk/blob/main/quickstart/quickstart.ipynb).

## Installation

To install GSCDK, run the following command:

For Unix-based platforms (like Linux or macOS):

```bash
python3 -m pip install gscdk
```

For Windows:

```bash
python -m pip install gscdk
```

### Upgrading GSCDK

If you've previously installed GSCDK and want to upgrade to the latest version, use the following command:

For Unix-based platforms:

```bash
python3 -m pip install --upgrade gscdk
```

For Windows:

```bash
python -m pip install --upgrade gscdk
```

## Building Go Smart Contracts Compiler

To build the `tinygo` command that supports building Go Smart Contracts, follow the instructions in [Building](./BUILDING.md). Once built, add the `tinygo` command to your PATH:

```bash
export PATH=$(pwd)/compiler/build:$PATH
```

## Using go-contract

`go-contract` is a powerful tool for managing your smart contract projects. Learn more about its features below:

### Initializing a Project

Use the "init" command to initialize a project with a specific contract name:

```bash
go-contract init mycontract
cd mycontract
```

### Generating ABI and Extra Code

Use the "gencode" command to generate ABI and extra code for smart contracts:

```bash
go-contract gencode
```

Note: Code generation is also the default option for the "build" command.

### Building Your Project

To compile the source code of your project, use the "build" command:

```bash
go-contract build
```

To disable code generation during the build process, use the `-gen-code=false` flag:

```bash
go-contract build -gen-code=false
```

To disable code optimization, use the `-d` or `--debug` option:

```bash
go-contract build -d
```

## Debugging

Before debugging, install `ipyeos`:

```bash
python3 -m pip install ipyeos
```

To update to a new version, use the following command:

```bash
python3 -m pip install -U ipyeos
```

Then run the debugging server:

```bash
eosdebugger
```

On Windows, use a Docker image to run a debugging server.

First, pull ipyeos docker image:

```bash
docker pull ghcr.io/uuosio/ipyeos:latest
```

then start a debugging server from docker:
```bash
docker run -it --rm -p 9090:9090 -p 9092:9092 -t ghcr.io/uuosio/ipyeos
```

Here's a sneak peek of the debugger in action:

![Debugging](https://github.com/uuosio/gscdk/blob/main/images/debugging.gif)

## Code Coverage Analysis

To generate a code coverage report in HTML, follow these steps:

1. Build your project:

```bash
go-contract build
```

2. Generate a coverage report:

```bash
TEST_COVERAGE=1 go test -coverprofile=coverage.out
```

3. Create an HTML report from the coverage data:

```bash
go tool cover -html=coverage.out
```

Here's an example of what the code coverage report looks like:

![Code Coverage](https://github.com/uuosio/gscdk/blob/main/images/code-coverage.png)

