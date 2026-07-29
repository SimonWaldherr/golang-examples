#!/bin/sh
set -eu

repository_root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
cd "$repository_root"

packages="
./beginner/initialized-pointer
./advanced/benchmark-loop
./advanced/generic-alias
./advanced/iterators
./advanced/structured-logging
./advanced/synctest
./advanced/waitgroup-go
./expert/flight-recorder
"

go test $packages
go vet $packages
go test ./advanced/benchmark-loop -run '^$' -bench '^BenchmarkJoin$' -benchtime=1x

unformatted=$(gofmt -l \
	beginner/initialized-pointer/main.go \
	advanced/benchmark-loop/join_test.go \
	advanced/generic-alias/main.go \
	advanced/iterators/main.go \
	advanced/structured-logging/main.go \
	advanced/synctest/timeout_test.go \
	advanced/waitgroup-go/main.go \
	expert/flight-recorder/main.go)

if [ -n "$unformatted" ]; then
	echo "These files need gofmt:"
	echo "$unformatted"
	exit 1
fi
