# `bump`
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/guilhem/bump)
[![bump](https://snapcraft.io/bump/badge.svg)](https://snapcraft.io/bump)

Command-line to bump version in a git repository

## Install

### [Snap](https://snapcraft.io/)

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/bump)

```sh
$ snap install bump
```

### [Homebrew](https://brew.sh/)

```sh
$ brew install guilhem/homebrew-tap/bump
```

### [Go get](https://golang.org/pkg/cmd/go/internal/get/)

```sh
$ go get github.com/guilhem/bump
```

## Usage

### Help

```sh
$ bump --help
Bump version

Usage:
  bump [command]

Available Commands:
  help        Help about any command
  major       Bump major version
  minor       Bump minor
  patch       Bump patch

Flags:
      --allow-dirty     allow usage of bump on dirty git
      --dry-run         Don't touch git repository
  -h, --help            help for bump
      --latest-tag      use latest tag, prompt tags if false (default true)
  -t, --toggle          Help message for toggle

Use "bump [command] --help" for more information about a command.
```

### Major

```sh
$ git tag
1.1.1
$ bump major
$ git tag
1.1.1
2.0.0
```

### Minor

```sh
$ git tag 
v1.1.1
$ bump minor
$ git tag
v1.1.1
v1.2.0
```

### Patch

```sh
$ git tag
1.1.1
$ bump patch
$ git tag
1.1.1
1.1.2
```
