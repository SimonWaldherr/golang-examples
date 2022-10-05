package main

import "fmt"

func factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}
func main() {
	fmt.Println(factorial(3))
}
