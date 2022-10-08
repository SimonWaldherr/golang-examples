package main

import (
	"fmt"
	"simonwaldherr.de/go/golibs/regex"
)

func main() {
	fmt.Println(regex.ReplaceAllString("FooBaR LoReM IpSuM", "\\W", ""))
	fmt.Println(regex.ReplaceAllString("FooBaR LoReM IpSuM", "[a-z]", ""))
	fmt.Println(regex.ReplaceAllString("FooBaR LoReM IpSuM", "[A-Z]", ""))
}
