package main

// import required modules
import (
	"fmt"
	"flag"
	"time"
	"strconv"
)

var x int64

func main() {
	start := time.Now()
	flag.Parse()
	s := flag.Arg(0)

	if s == "" {
		s = "10"
	}

	x, err := strconv.ParseInt(s, 10, 0)

	if err != nil {
		fmt.Println(err)
		x = 10
	}

	fibonacci(x)
	elapsed := time.Since(start)
	fmt.Println("\ntime:", elapsed)
}

func fibonacci(n int64) {
	var a int64 = 0
	var b int64 = 1
	var i int64
	var sum int64
	for i = 0; i < n; i++ {
		fmt.Println(a)
		sum = a + b
		a = b
		b = sum
	}
}
