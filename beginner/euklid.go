package main

import (
	"fmt"
	"os"
	"strconv"
)

func euklid(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	if len(os.Args) > 2 {
		ia1, _ := strconv.ParseInt(os.Args[1], 10, 0)
		ia2, _ := strconv.ParseInt(os.Args[2], 10, 0)
		fmt.Println(euklid(ia1, ia2))
	}
}
