// Description: Channel example 2
// Tags: channel, range, for, loop, close, close channel, close ch
package main

import "fmt"

func main() {
	numbersArray := []int{1, 2, 3, 4, 5, 6, 7, 8}
	channel := make(chan int)

	go sumNumbers(numbersArray[:len(numbersArray)/2], channel)
	go sumNumbers(numbersArray[len(numbersArray)/2:], channel)

	number1 := <-channel
	number2 := <-channel

	fmt.Println("Sum of two numbers in first goroutine: ", number1)
	fmt.Println("Sum of two numbers in second goroutine: ", number2)

	fmt.Println("Total : ", number1+number2)
}

func sumNumbers(numbers []int, channel chan int) {
	result := 0
	for _, value := range numbers {
		result += value
	}
	channel <- result
}
