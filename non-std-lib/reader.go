package main

import (
	"bufio"
	"log"

	foobar "github.com/SimonWaldherr/golang-examples/non-std-lib/packages"
)

func main() {
	M := foobar.NewReader()

	bytes := make([]byte, 255)

	bufio.NewReader(M).Read(bytes)

	log.Printf("%s", bytes)
}
