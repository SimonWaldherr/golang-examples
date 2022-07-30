package main

import "fmt"

func main() {
	queue := make(chan int, 6)
	queue <- 23
	queue <- 42
	queue <- 1337
	queue <- 9999
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
