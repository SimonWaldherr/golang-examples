package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a temporary file
	// NOTE:
	// the permisson for this file will be 0600
	// if you want to make it readable for other user
	// change by f.Chmod(0o644)
	f, err := os.CreateTemp("", "sample")
	if err != nil {
		fmt.Println(err)
		return
	}

	// demo for chage file mode
	// err = f.Chmod(0o644)

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
