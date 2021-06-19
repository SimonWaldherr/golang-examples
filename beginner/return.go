package main

// import required modules
import (
	"fmt"
)

func named() (str string) {
	str = "lorem"
	return
}

func typed() string {
	var str = "ipsum"
	return str
}

// main function
func main() {
	fmt.Println(named())
	fmt.Println(typed())
}
