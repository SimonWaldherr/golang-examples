package main

// import required modules
import (
	"fmt"
	"os"
)

// main function
func main() {

	// print each argument
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
