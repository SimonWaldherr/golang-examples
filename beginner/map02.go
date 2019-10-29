package main

import "fmt"

func main() {
	createMap()
}

func createMap() {
	//initialize map with key and value as string
	var stringMap = make(map[string]string)

	//add keys to map
	stringMap["A"] = "AAA"
	stringMap["B"] = "BBB"
	stringMap["C"] = "CCC"

	fmt.Print(stringMap)

	//delete key
	delete(stringMap, "A")

	//initialize map with key and value using a map literal
	var intMap = map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	fmt.Print(intMap)

	//access item from map
	value, ok := intMap[1] //ok is true if item exists
	if ok {
		fmt.Printf("Key = 1; Value =%v", value)
	}
}
