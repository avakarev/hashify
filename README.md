# hashify

[![Latest Release](https://img.shields.io/github/release/avakarev/hashify.svg)](https://github.com/avakarev/hashify/releases)
[![CI Status](https://github.com/avakarev/hashify/actions/workflows/ci.yml/badge.svg)](https://github.com/avakarev/hashify/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/avakarev/hashify)](https://goreportcard.com/report/github.com/avakarev/hashify)

CLI utility to calculate file checksums

## Installation

#### macOS

With [Homebrew](https://brew.sh/): `brew tap avakarev/tap && brew install hashify`

### Binaries

[Binaries](https://github.com/avakarev/hashify/releases) for Linux and macOS

### From source

Make sure you have a working Go environment (Go 1.12 or higher is required).
See the [install instructions](http://golang.org/doc/install.html).

Compiling `hashify` is easy, simply run:

    git clone https://github.com/avakarev/hashify.git
    cd hashify
    make build

## Usage

    ○ $:> hashify -md5 -sha256 go.mod go.sum test/file
    go.mod
       md5 | 88e79e82744fd709ca1c5918f88fe9e9
    sha256 | 03f3b24637f943604160573177815de6067c89c1993e38ea4d39595db246d51c

    go.sum
       md5 | 9073179bce05f8762267b2017ecd8e5c
    sha256 | 270473f17ead3f907fda3372bfc3a4fed2a45f3527af79cb5d84fbec9c50cfbf

    test/file
       md5 | 8ddd8be4b179a529afa5f2ffae4b9858
    sha256 | 03ba204e50d126e4674c005e04d82e84c21366780af1f43bd54a37816b6ab340

## License

`hashify` is licensed under MIT license. (see [LICENSE](./LICENSE))
