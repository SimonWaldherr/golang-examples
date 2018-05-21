package main

import (
	"fmt"
	"io/ioutil"
	mmd "simonwaldherr.de/go/micromarkdownGo"
)

func main() {
	mdfile, _ := ioutil.ReadFile("./demo.md")

	md := mmd.Micromarkdown(string(mdfile))

	fmt.Println(string(md))
}
