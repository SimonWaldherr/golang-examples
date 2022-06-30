#!/bin/bash 

# build/combile wasm.go to wasm-file
echo "build"
GOOS=js GOARCH=wasm go build -o cc.wasm wasm.go

# copy wasm_exec.js to this dir
echo "copy wasm_exec.js"
cp $(go env GOROOT)/misc/wasm/wasm_exec.js wasm_exec.js

# run http server
echo "run server at http://localhost:8080/"
open "http://localhost:8080/"
go run server.go 
