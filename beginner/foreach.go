package main

import (
	"fmt"
)

// declare variables and define array content
var strarray = []string{"lorem", "ipsum", "dolor", "sit", "amet"}

func main() {
	for index, data := range strarray {
		fmt.Println(index, data)
	}
}
