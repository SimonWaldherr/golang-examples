package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	gofile, err1 := ioutil.ReadFile("./file.go")
	check(err1)
	output := []byte(gofile)
	fmt.Println(string(gofile))
	err2 := ioutil.WriteFile("./file.txt", output, 0644)
	check(err2)
}
