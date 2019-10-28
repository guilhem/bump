# `bump`

Command-line to bump version in a git repository

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