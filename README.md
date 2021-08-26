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

This will build tinygo command in compiler/build directory that support for build Go Smart Contracts.

## Add tinygo to PATH environment variable
```
export PATH=$(pwd)/compiler/build:$PATH
```

# Install from Release Binary

First download binary file from [releases](https://github.com/uuosio/uuosio.gscdk/releases)

For tar.gz file

```bash
tar -C /usr/local -xzf release.tar.gz
export PATH=/usr/local/uuosio.gscdk/bin:$PATH
```

Install debian package directly on Ubuntu platform
```bash
sudo apt install ./release.deb
```

# Examples

# Tutorials

# Additional Options Other than Official Tinygo Commands

## gen-code

If gen-code option is true, generate code for Go Smart Contracts or not. Default to "true". Only has effect on eosio target.

## strip

If strip option is true, strip custom section for wasm file to reduce size. Default to "true". Only has effect on eosio target

