// Description: Embedding files into the binary
// Tags: embed, embed file, embed folder, embed files, embed folder, embed files, embed folder, embed
package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed example.csv
var fileString string

//go:embed packages/*.go
var folder embed.FS

func main() {
	fmt.Println("example.csv:")
	fmt.Println(fileString)

	files, err := folder.ReadDir("packages")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nEmbedded files:")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	content, _ := folder.ReadFile("packages/foobar.go")
	fmt.Println("\npackages/foobar.go:")
	fmt.Println(string(content))
}
