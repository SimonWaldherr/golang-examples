package main

import (
	"fmt"
	"os"

	mmd "simonwaldherr.de/go/micromarkdownGo"
)

func main() {
	mdfile, _ := os.ReadFile("./demo.md")

	md := mmd.Micromarkdown(string(mdfile))

	fmt.Println(string(md))
}
