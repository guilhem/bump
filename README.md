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

### Go (1.27.1 or newer)

```sh
$ go install github.com/guilhem/bump@latest
```

The executable is installed in `$(go env GOPATH)/bin`; add that directory to
`PATH` if needed. Prebuilt archives are available from [GitHub releases](https://github.com/guilhem/bump/releases).

## Getting started

Run `bump` inside a clean Git repository with at least one lightweight SemVer tag
in the current branch's history:

```sh
git tag v0.1.0                  # only when starting a version history
bump patch --dry-run            # preview the next version
bump patch                      # create v0.1.1 on HEAD
git push origin v0.1.1           # publish the tag when ready
```

The `v` prefix is preserved. `--latest-tag=false` selects a previous tag
interactively; `--allow-dirty` permits an uncommitted worktree. `bump` creates
local lightweight tags and does not push them. Command failures return a nonzero
exit code; `bump --help` also works outside a repository.

## Usage

### Help

```sh
$ bump --help
Bump version

Usage:
  bump [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  major       Bump major version
  minor       Bump minor
  patch       Bump patch

Flags:
      --allow-dirty   allow usage of bump on dirty git
      --dry-run       Don't touch git repository
  -h, --help          help for bump
      --latest-tag    use latest tag, prompt tags if false (default true)
  -t, --toggle        Help message for toggle

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

## Development

With Go 1.27.1 or newer:

```sh
go test ./...
go vet ./...
sh scripts/smoke-test.sh
go build .
```

CI runs the tests and a GoReleaser 2.18.2 snapshot on pushes and pull requests.
To reproduce the archive generation locally:

```sh
goreleaser release --snapshot --clean --skip=snapcraft
```

Tagged pushes publish GitHub archives, a Homebrew formula, and a `core24` Snap.
Maintainers must configure `TAP_TOKEN` for `guilhem/homebrew-tap` and
`SNAPCRAFT_TOKEN` with Snap Store credentials. The release job grants its
`GITHUB_TOKEN` permission to write release assets. GoReleaser still supports the
existing Homebrew formula configuration but warns that `brews` is deprecated;
a move to casks is a separate distribution change.
