## About

offvsix is a command-line utility written in Go that can
download Visual Studio Code extensions (.vsix files) directly
from the Visual Studio Marketplace. This project is inspired by
the Python-based [offvsix](https://github.com/gni/offvsix) tool
of the same name but this copy is rewritten in Go.

## Features

* Download specific versions of VS Code extensions
* Single binary with no dependencies
* Cross-platform support: Windows, macOS, Linux, and BSD
* **Bulk download:** Supply a text file with a list of extensions to download them all at once!

## Installation

You can install the binary using `go install`:

```bash
go install github.com/0x1eef/offvsix/cmd/offvsix@latest
~/go/bin/offvsix ms-python.python
```

Or build from source:

```bash
git clone https://github.com/0x1eef/offvsix.git
cd offvsix
make build
./bin/offvsix ms-python.python
```

## Usage

### Basics

```bash
offvsix <publisher.extension>
```

For example:

```bash
offvsix ms-python.python
```

### Bulk

To download multiple extensions, you can use a text file where each line is an extension name:

```bash
offvsix -f extensions.txt
```

### Install

```bash
code-server --install-extension golang.go-0.51.1.vsix
```

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
