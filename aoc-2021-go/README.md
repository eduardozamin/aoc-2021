# Advent of Code 2021 - Go

Solutions for the Advent of Code 2021 implemented in Go.

It is implemented using Cobra (https://github.com/spf13/cobra) as CLI framework and
follows the project layout recommendation from https://github.com/golang-standards/project-layout


## Environment Setup

```
# install global dependencies
scripts/setup-dev.sh

# run tests
go test -v ./...

# create binary output
go build

# run
./aoc-2021 --help
```

## Executing the solutions

The solutions are called as Cobra commands (available in `cmd/`) and implemented as libraries
in `internal/pkg/`.

The input files are expected to be found at `input/input-day<XX>`

To run the solution for day04 for example you have to run:
```
# using "go run"
go run main.go solve day04

# or via build + execution
go build
./aoc-2021 solve day04
```

The results are displayed in stdout.
