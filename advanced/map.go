package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string)

	//Insertionn in a map
	myMap["A"] = "apple"
	myMap["B"] = "ball"
	myMap["C"] = "cat"

	//Traverse a map
	fmt.Println("Printing Map..........")
	for key, value := range myMap {
		fmt.Println("Key: ", key, "  Value: ", value)
	}

	//Delete a key-value pair from a map
	//Deleting Key: "C" from map
	fmt.Println("Deleting a key from Map..........")
	delete(myMap, "C")

	//Find the value associated with a key
	fmt.Println("Finding a key from Map..........")
	fmt.Println("Value for key 'A': ", myMap["A"])
}
