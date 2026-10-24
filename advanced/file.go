// Description: This program shows how to read and write files in Go
// Tags: file, read, write, file read, file write, file read wr
package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := "./file.go"

	// os.Stat returns file info
	finfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(filename + " doesn't exist")
	} else {
		// IsDir returns boolean true if finfo is a directory
		if finfo.IsDir() {
			fmt.Println(filename + " is a directory")
		} else {

			// ReadFile returns the input of the file
			gofile, err1 := os.ReadFile(filename)
			check(err1)
			output := []byte(gofile)
			fmt.Println(string(gofile))

			// WriteFile writes data to a file
			err2 := os.WriteFile("./file.txt", output, 0o644)
			check(err2)
		}
	}
}
