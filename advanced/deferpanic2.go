package main

import (
	"fmt"
)

func main() {
	stop := 10
	for i := 0; i < 15; i++ {
		func() {
			defer func() {
				recover()
			}()

			if i == 10 {
				fmt.Println("i10")
				panic(0)
			}
			if i == 14 {
				fmt.Println("i14")
				i = 0
				stop--
				panic(0)
			}
		}()
		if stop == 0 {
			break
		}
	}
}
