#!/bin/bash

export GOMAXPROCS=8
time go run parallel.go true
time go run parallel.go false
