package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := "./file.go"

	//os.Stat returns file info
	finfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(filename + " doesn't exist")
	} else {

		//IsDir returns boolean true if finfo is a directory
		if finfo.IsDir() {
			fmt.Println(filename + " is a directory")
		} else {

			//ReadFile returns the input of the file
			gofile, err1 := ioutil.ReadFile(filename)
			check(err1)
			output := []byte(gofile)
			fmt.Println(string(gofile))

			//WriteFile writes data to a file
			err2 := ioutil.WriteFile("./file.txt", output, 0644)
			check(err2)
		}
	}
}
