# `bump`
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/guilhem/bump)
[![bump](https://snapcraft.io/bump/badge.svg)](https://snapcraft.io/bump)

Command-line to bump version in a git repository

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/bump)

## Usage

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
1.1.1
$ bump minor
$ git tag
1.1.1
1.2.0
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