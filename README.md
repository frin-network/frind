
Frind
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/frin-network/frind/)

Frind is the reference full node frin implementation written in Go (golang).

## What is Frin

Frin is a fork of Kaspa with an ASIC resistance implementation
Kaspa is an attempt at a proof-of-work cryptocurrency with instant confirmations and sub-second block times. It is based on [the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a generalization of Nakamoto consensus.

## Requirements

Go 1.18 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install frind including all dependencies:

```bash
$ git clone https://github.com/frin-network/frind/
$ cd frind
$ go install . ./cmd/...
```

- frind (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.


## Getting Started

frind has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ frind
```

## Discord
Join our discord server using the following link: https://discord.gg/

## Issue Tracker

The [integrated github issue tracker](https://github.com/frin-network/frind/issues)
is used for this project.


## Documentation

The [documentation](https://github.com/frin-network/docs) is a work-in-progress

## License

frind is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
