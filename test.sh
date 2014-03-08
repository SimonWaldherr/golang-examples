#!/bin/bash

go run HelloWorld.go

go run var.go

go run array.go

go run for.go

go run args.go string string2

go run input.go

go run pythagoras.go 10 15 ?

go run time.go

go run random.go

go run cat.go

go run if.go Hello

go run modulo.go

go run lifo.go

go run split.go

go run splitbyregex.go

go run fibonacci.go

go run prime.go 32

go run json.go

go run httpd.go

go run proxy.go

go run exec.go

export GOMAXPROCS=8
time go run parallel.go true
time go run parallel.go false
