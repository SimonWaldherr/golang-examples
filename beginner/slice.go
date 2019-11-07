package main

// import required modules
import (
	"fmt"
)

// main function
func main() {
	var str string = "Lorem ipsum dolor sit amet"
	fmt.Println(str[6:11])

	//Make a new slice using make
	s := make([]string, 3)

	//Appending elements to existing slice
	s = append(s, "abc")
	fmt.Println(s)
}
