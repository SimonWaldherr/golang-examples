// This is a simple program that shows how to use go in taking inputs from a console... It takes in name and age as inputs and then uses conditionals to see if the age is permitted to continue
package main

import "fmt"

func main() {
	fmt.Printf("What's your name? : ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("How old are you? : ")
	var age int
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println(name, "You are too young to continue! ")

	} else {
		fmt.Println("Nice to see you", name)
	}
}
