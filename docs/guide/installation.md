# Install go-bitsong

This guide will explain how to install the `bitsongd` and `bitsongcli` entrypoints onto your system. With these installed on a server, you can participate in the testnet as either a [Full Node](./join-testnet.md) or a [Validator](../validators/validator-setup.md).

## Install Go

Install `go` by following the [official docs](https://golang.org/doc/install). Remember to set your `$PATH` environment variable, for example:

```bash
mkdir -p $HOME/go/bin
echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bash_profile
source ~/.bash_profile
```

::: tip
**Go 1.13+** is required for `go-bitsong`.
:::

## Install the binaries

Next, let's install the latest version of `go-bitsong`. Make sure you `git checkout` the correct [released version](https://github.com/bitsongofficial/go-bitsong/releases).

```bash
git clone -b <latest-release-tag> https://github.com/bitsongofficial/go-bitsong
cd go-bitsong && make install
```

If this command fails due to the following error message, you might have already set `LDFLAGS` prior to running this step.

```
# github.com/bitsongoffcial/go-bitsong/cmd/bitsongd
flag provided but not defined: -L
usage: link [options] main.o
...
make: *** [install] Error 2
```

Unset this environment variable and try again.

```
LDFLAGS="" make install
```

> _NOTE_: If you still have issues at this step, please check that you have the latest stable version of GO installed.

That will install the `bitsongd` and `bitsongcli` binaries. Verify that everything is OK:

```bash
$ bitsongd version --long
$ bitsongcli version --long
```

`bitsongcli` for instance should output something similar to:

```shell
name: go-bitsong
server_name: bitsongd
client_name: bitsongcli
version: 0.3.0
commit: 1dc65dbb84309435ab9e87c67202c071e0f0745c
build_tags: netgo,ledger
go: go version go1.13.6 linux/amd64
```

### Build Tags

Build tags indicate special features that have been enabled in the binary.

| Build Tag | Description                                     |
| --------- | ----------------------------------------------- |
| netgo     | Name resolution will use pure Go code           |
| ledger    | Ledger devices are supported (hardware wallets) |

## Next

Now you can [join the public testnet](./join-testnet.md) or [create you own testnet](./deploy-testnet.md)
