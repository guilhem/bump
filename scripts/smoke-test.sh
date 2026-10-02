#!/bin/sh
set -eu

cd "$(dirname "$0")/.."
fixture=$(mktemp -d)
trap 'rm -rf "$fixture"' EXIT HUP INT TERM
go build -o "$fixture/bump" .
cd "$fixture"
./bump --help >/dev/null
git init -q repo
cd repo
git -c user.name=Test -c user.email=test@example.com -c commit.gpgsign=false commit -q --allow-empty -m initial
git tag nightly
if ../bump patch >"$fixture/error" 2>&1; then
  echo "expected an error for non-version tags" >&2
  exit 1
fi
if grep -q panic "$fixture/error"; then
  cat "$fixture/error" >&2
  exit 1
fi
git tag v1.2.3
git tag v1.3.0-rc.1
../bump patch --dry-run
if git rev-parse --verify refs/tags/v1.3.0 >/dev/null 2>&1; then
  echo "dry-run created a tag" >&2
  exit 1
fi
../bump patch
git rev-parse --verify refs/tags/v1.3.0 >/dev/null
printf dirty > untracked
if ../bump minor >"$fixture/error" 2>&1; then
  echo "expected an error for a dirty worktree" >&2
  exit 1
fi
../bump minor --allow-dirty
git rev-parse --verify refs/tags/v1.4.0 >/dev/null
rm untracked
git tag -d v1.2.3 v1.3.0-rc.1 v1.3.0 v1.4.0 >/dev/null
git tag 2.0.0
../bump major
git rev-parse --verify refs/tags/3.0.0 >/dev/null
echo "CLI smoke passed: help, invalid tags, prerelease, dry-run, dirty worktree, tag prefixes"
