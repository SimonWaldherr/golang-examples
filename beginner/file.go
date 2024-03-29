package main

import "os"

// OpenCreateFile opens an existing file or creates a new file if it does not exist
func OpenCreateFile(path string) *os.File {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		return f
	}
	// open file in read-write mode
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

func main() {
	fp := OpenCreateFile("test.txt")
	fp.WriteString("Hello World")
	fp.Close()
}
