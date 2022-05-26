#!/bin/sh
set -e -x

go build ./...
go build -o ./server ./cmd/server
go build -o ~/.terraform.d/plugins/ransford.org/edu/foo/0.3.1/darwin_arm64/terraform-provider-foo ./cmd/provider
