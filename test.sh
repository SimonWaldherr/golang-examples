#!/bin/bash

function killgop {
  echo "killing $1"
  if [[ "$OSTYPE" == "linux-gnu" ]]; then
    kill -9 $1
  elif [[ "$OSTYPE" == "darwin"* ]]; then
    kill -TERM $1
  fi
}

echo "beginner"
echo "HelloWorld"
go run ./beginner/HelloWorld.go

echo "var"
go run ./beginner/var.go

echo "if"
go run ./beginner/if.go Hello

echo "array"
go run ./beginner/array.go

echo "function"
go run ./beginner/function.go

echo "for"
go run ./beginner/for.go

echo "foreach"
go run ./beginner/foreach.go

echo "args"
go run ./beginner/args.go string string2

echo "input"
go run ./beginner/input.go

echo "flag"
go run ./beginner/flag.go

echo "dir"
go run ./beginner/dir.go

echo "time"
go run ./beginner/time.go

echo "random"
go run ./beginner/random.go

echo "cat"
go run ./beginner/cat.go

echo "modulo"
go run ./beginner/modulo.go

echo "split"
go run ./beginner/split.go

echo "hashing"
go run ./beginner/hashing.go

echo "ackermann"
go run ./beginner/ackermann.go

echo "euklid"
go run ./beginner/euklid.go

echo "variadic function"
go run ./beginner/functionvariadic.go

echo "delete fom slice"
go run ./beginner/deleteFromSlice.go

echo "advanced"
echo "pythagoras"
go run ./advanced/pythagoras.go 10 15 ?

echo "wait"
go run ./advanced/wait.go

echo "lifo"
go run ./advanced/lifo.go

echo "regex"
go run ./advanced/regex.go

echo "fibonacci"
go run ./advanced/fibonacci.go

echo "prime"
go run ./advanced/prime.go 32

echo "numbers"
go run ./advanced/numbers.go

echo "json"
go run ./advanced/json.go

echo "exec"
go run ./advanced/exec.go

echo "suicide"
go run ./advanced/suicide.go

echo "in array search"
go run ./advanced/in_array.go

echo "benchmarking json marshal and unmarshal"
go test -bench=. -benchmem ./advanced/json_bench/main_test.go


echo "expert"
echo "color"
go run ./expert/color.go

echo "telnet"
go run ./expert/telnet.go &
sleep 10
nc localhost 2223 < lorem
killgop $!

echo "httpd"
go run ./expert/httpd.go &
sleep 10
for i in {1..10}; do curl http://localhost:8080; done;
killgop $!

echo "httpsd"
cd expert
go run ./httpsd.go &
sleep 5
echo "waiting 10 seconds - then make 10 connections"
sleep 10
for i in {1..10}; do curl https://localhost:4443 --insecure; done;
killgop $!

cd ..

echo "cookies"
go run ./expert/cookies.go &
sleep 10
for i in {1..2}; do curl http://localhost:8080; done;
killgop $!

echo "proxy"
go run ./expert/proxy.go &
sleep 10
curl https://localhost:8080 --insecure
killgop $!

echo "ppk-crypto"
go run ./export/ppk-crypto.go

echo "image"
cd expert
go run ./image.go

echo "sqlite"
go get github.com/mxk/go-sqlite/sqlite3
go run ./sqlite.go insert test
go run ./sqlite.go select

cd ..

echo "parallel"
export GOMAXPROCS=8
time go run ./expert/parallel.go true
time go run ./expert/parallel.go false

echo "dynparallel"
time go run ./expert/dynparallel.go 8
