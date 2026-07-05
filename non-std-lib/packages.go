package main

import (
	"fmt" //import a standard package

	foobar "github.com/SimonWaldherr/golang-examples/non-std-lib/packages" //import a "local" package folder (package name: foobar)
	"simonwaldherr.de/go/golibs/as"                                        //the path to the package
	convert_to "simonwaldherr.de/go/golibs/as"                             //with a different name
)

func main() {
	fmt.Println(as.String(0x23))           //access the package with the last folder name (as)
	fmt.Println(convert_to.String(0xfefe)) //or with a self-selected name
	fmt.Println(foobar.PublicFunc(5))      //use the functions from ./packages/foobar.go
}
