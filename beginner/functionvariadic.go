package main

import "fmt"

func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sum(2, 3, 5, 7, 11, 13, 17, 19, 23, 29))

	numbers := []int{31, 37, 41, 43, 47, 53, 59}

	fmt.Println(sum(numbers...))
}
