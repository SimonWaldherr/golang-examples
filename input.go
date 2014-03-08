package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Println("Enter a number")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(i)
	}
}
