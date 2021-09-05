Go Smart Contracts Development Kit

# Quick Start

<h3>
  <a
    target="_blank"
    href="https://mybinder.org/v2/gh/uuosio/uuosio.gscdk/main?filepath=quickstart/quickstart.ipynb"
  >
    Quick Start
    <img alt="Binder" valign="bottom" height="25px"
    src="https://mybinder.org/badge_logo.svg"
    />
  </a>
</h3>

# Build Go Smart Contracts Compiler

Follow the steps in [Building](./BUILDING.md)

That will build tinygo command in compiler/build directory that support build Go Smart Contracts.

#### Set PATH

```bash
export PATH=$(pwd)/compiler/build:$PATH
```

# Install from Pre-built Binary

First download binary file from [releases](https://github.com/uuosio/uuosio.gscdk/releases)

For tar.gz file

```bash
tar -C /usr/local -xzf uuosio.gscdk-linux-0.1.0.tar.gz
export PATH=/usr/local/uuosio.gscdk/bin:$PATH
```

Install debian package directly on Ubuntu platform

```bash
sudo apt install ./uuosio.gscdk-linux-0.1.0.deb
```

# eosio-go Command

All commands based on examples/hello example

First cd to hello directory:

```bash
cd examples/hello
```

#### Building

```bash
eosio-go build -o hello.wasm .
```

#### Generating ABI and Extra Code for Smart Contracts

```
eosio-go gencode
```

Code generation is also the default option for "build" command


#### Disable Code Generation during Building

```bash
eosio-go build -gen-code=false -o hello .
```

#### Disable Striping WASM File after Building

```bash
eosio-go build -strip=false -o hello .
```
