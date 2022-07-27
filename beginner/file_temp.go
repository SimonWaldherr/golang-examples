package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a temporary file
	f, err := os.CreateTemp("", "sample")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the name of the just created temporary file
	fmt.Println("Temp file name:", f.Name())

	// Clean up the file after we're done
	defer os.Remove(f.Name())

	// Write some data to the file
	_, err = f.Write([]byte("Hello\nWorld\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
