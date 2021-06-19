package main

// import required modules
import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

var x uint64

func main() {
	start := time.Now()
	flag.Parse()
	s := flag.Arg(0)

	if s == "" {
		s = "140"
	}

	x, err := strconv.ParseUint(s, 10, 0)

	if err != nil {
		fmt.Println(err)
		x = 40
	}

	fibonacci(x)
	elapsed := time.Since(start)
	fmt.Println("\ntime:", elapsed)
}

func fibonacci(n uint64) {
	var a uint64 = 0
	var b uint64 = 1
	var i uint64
	var sum uint64
	for i = 0; i < n; i++ {
		fmt.Println(a)
		sum = a + b
		a = b
		b = sum
	}
}
