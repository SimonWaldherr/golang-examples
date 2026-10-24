package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err == nil {
		fmt.Println(wd)
		files, _ := os.ReadDir(wd)
		for _, f := range files {
			fmt.Printf("-> %v\n", f.Name())
		}
	}
}
