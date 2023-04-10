// Description: Anonymous functions and goroutines
// Tags: anonymous, function, anonymous function, anonymous function goroutine, anonymous function goroutine, goroutine, goroutine anonymous function, goroutine anonymous function, anonymous function, anonymous function goroutine, anonymous function goroutine, goroutine, goroutine anonymous function, goroutine anonymous function
package main

import (
	"fmt"
)

func main() {
	stop := make(chan bool)

	go func() {
		fmt.Println("this is Println inside an anonymous goroutine")
		stop <- true
	}()
	func() {
		fmt.Println("this is Println inside an anonymous function")
	}()
	<-stop
}
