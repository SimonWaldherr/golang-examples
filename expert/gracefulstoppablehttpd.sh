#!/bin/bash

go run gracefulstoppablehttpd.go &
sleep 10

for i in {1..10}; do curl http://localhost:8080 --insecure & done;

sleep 1

for i in {1..10}; do curl http://localhost:8080 --insecure & done;

go run gracefulstoppablehttpd.go &

sleep 1

curl http://localhost:8080/stop --insecure 

for i in {1..10}; do curl http://localhost:8080 --insecure & done;

curl http://localhost:8080/stop --insecure
