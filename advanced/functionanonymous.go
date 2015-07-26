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
