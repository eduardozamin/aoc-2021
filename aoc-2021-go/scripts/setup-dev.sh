#!/usr/bin/env bash


# bash strict settings
set -eou pipefail

set -x
export GOPROXY="https://proxy.golang.org"
go get -u golang.org/x/tools/cmd/goimports
go get -u github.com/spf13/cobra
go get github.com/spf13/cobra/cobra/cmd
go install github.com/spf13/cobra/cobra
go mod tidy
