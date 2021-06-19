package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func q(a float64) float64 {
	return a * a
}

func main() {
	if len(os.Args) == 4 {
		switch "?" {
		case os.Args[1]:
			b, _ := strconv.ParseFloat(os.Args[2], 0)
			c, _ := strconv.ParseFloat(os.Args[3], 0)
			fmt.Println(math.Sqrt(q(float64(c)) - q(float64(b))))
		case os.Args[2]:
			a, _ := strconv.ParseFloat(os.Args[1], 0)
			c, _ := strconv.ParseFloat(os.Args[3], 0)
			fmt.Println(math.Sqrt(q(float64(c)) - q(float64(a))))
		case os.Args[3]:
			a, _ := strconv.ParseFloat(os.Args[1], 0)
			b, _ := strconv.ParseFloat(os.Args[2], 0)
			fmt.Println(math.Sqrt(q(float64(a)) + q(float64(b))))
		default:
			fmt.Println("one of the arguments have to be a \"?\"")
		}
	} else {
		fmt.Println("this app needs 3 arguments")
	}
}
