package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	log.Println("first log output")
	log.Printf("second log output\n")
}
