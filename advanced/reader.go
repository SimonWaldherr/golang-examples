package main

import (
	"./packages/"
	"bufio"
	"log"
)

func main() {
	M := foobar.NewReader()

	bytes := make([]byte, 255)

	bufio.NewReader(M).Read(bytes)

	log.Printf("%s", bytes)
}
