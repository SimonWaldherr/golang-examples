// Description: Structs in Go
// Tags: struct
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a new person
	p := Person{Name: "John", Age: 42}

	// Print the person
	fmt.Printf("%v, %T, %#v\n", p, p, p)
	// Output: {John 42}, main.Person, main.Person{Name:"John", Age:42}

	// Print the person's name
	fmt.Printf("%v, %T, %#v\n", p.Name, p.Name, p.Name)
	// Output: John, string, "John"

	// Print the person's age
	fmt.Printf("%v, %T, %#v\n", p.Age, p.Age, p.Age)
	// Output: 42, int, 42
}
