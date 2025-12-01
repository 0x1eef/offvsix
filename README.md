## About

offvsix is a command-line utility written in Go that can
download Visual Studio Code extensions (.vsix files) directly
from the Visual Studio Marketplace. This project is inspired by
the Python-based [@gni/offvsix](https://github.com/gni/offvsix) tool
of the same name but this copy is written in Go.

## Features

* Download specific versions of VS Code extensions
* Supply a text file with a list of extensions to download them all at once
* Single binary with no dependencies
* Cross-platform support: Windows, macOS, Linux, and BSD

## Install

#### Binaries

Prebuilt binaries are made available via [GItHub actions](https://github.com/0x1eef/offvsix/actions/runs/19806477120):

* [offvsix-windows-amd64 (v0.2.0)](https://github.com/0x1eef/offvsix/actions/runs/19806477120/artifacts/4718983018)
* [offvsix-linux-amd64 (v0.2.0)](https://github.com/0x1eef/offvsix/actions/runs/19806477120/artifacts/4718980945)
* [offvsix-darwin-amd64 (v0.2.0)](https://github.com/0x1eef/offvsix/actions/runs/19806477120/artifacts/4718980702)
* [offvsix-freebsd-amd64 (v0.2.0)](https://github.com/0x1eef/offvsix/actions/runs/19806477120/artifacts/4718985996)

#### Package

    user@localhost$ go install github.com/0x1eef/offvsix/cmd/offvsix@latest
    user@localhost$ ~/go/bin/offvsix golang.go

#### Source

    user@localhost$ git clone https://github.com/0x1eef/offvsix.git
    user@localhost$ cd offvsix
    user@localhost$ make build
    user@localhost$ ./bin/offvsix golang.go

## Usage

#### Basics

The following examples demonstrate how to download individual
extensions, and by a specific version:

    user@localhost$ offvsix <publisher.extension>
    user@localhost$ offvsix golang.go
    user@localhost$ offvsix -v 0.50.0 golang.go

#### Bulk

The following example assumes extensions.txt contains a list of
extensions, with one extension per line:

    user@localhost$ offvsix -f extensions.txt

#### Install

The following example demonstrates how to install an extension
with [code-server](https://github.com/coder/code-server):

    user@localhost$ code-server --install-extension golang.go-0.51.1.vsix

## Credit

This project is a port of the Python-based [@gni/offvsix](https://github.com/gni/offvsix)
written by [@gni](https://github.com/gni). I studied the source code of
[@gni/offvsix](https://github.com/gni/offvsix)
to first understand how to implement the same functionality
in Go, and their work made this project possible.

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
