package main

import (
	"fmt"
	"log"
	"os"
)

func write() {
	content := []byte("Hello, Golang!")
	err := os.WriteFile("example.txt", content, 0o644)
	if err != nil {
		log.Fatal(err)
	}
}

func read() {
	content, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File content: %s\n", content)
}

func main() {
	write()
	read()
}
