package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func write() {
	content := []byte("Hello, Golang!")
	err := ioutil.WriteFile("example.txt", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func read() {
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File content: %s\n", content)
}

func main() {
	write()
	read()
}
