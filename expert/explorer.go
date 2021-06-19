package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err == nil {
		fmt.Println(wd)
		files, _ := ioutil.ReadDir(wd)
		for _, f := range files {
			fmt.Printf("-> %v\n", f.Name())
		}
	}
}
