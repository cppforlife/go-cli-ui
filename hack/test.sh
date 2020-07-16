#!/bin/bash

set -e

if [ -z "$GITHUB_ACTION" ]; then
  go clean -testcache
fi

set -u

go fmt ./...
go mod vendor
go mod tidy

go test -test.v ./... $@

go build ./ui/
go build ./errors/
go build -o /dev/null github.com/cppforlife/go-cli-ui/examples/...

echo SUCCESS
