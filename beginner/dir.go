package main

import (
	"os"
	"fmt"
	"strings"
)


func main() {
	dir, _ := os.Getwd()
	fmt.Println(strings.Replace(dir," ","\\ ",-1))
}
