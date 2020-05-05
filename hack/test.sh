#!/bin/bash

set -e

if [ -z "$GITHUB_ACTION" ]; then
  go clean -testcache
fi

set -u

go fmt ./...
go test ./...

go build ./ui/
go build -o /dev/null github.com/cppforlife/go-cli-ui/examples/...

echo SUCCESS
