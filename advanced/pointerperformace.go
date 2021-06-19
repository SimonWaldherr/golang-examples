package main

import (
	"log"
	"time"
)

func plusP(in *int) {
	*in = *in + 1
}

func plusPpp(in *int) {
	*in++
}

func plus(in int) int {
	return in + 1
}

func main() {
	var value int
	var start time.Time
	var elapsed time.Duration

	value = 5

	start = time.Now()

	for i := 0; i < 1600000; i++ {
		plusP(&value)
	}

	elapsed = time.Since(start)
	log.Printf("calculated %v with pointer took: %s", value, elapsed)

	value = 5

	start = time.Now()

	for i := 0; i < 1600000; i++ {
		plusPpp(&value)
	}

	elapsed = time.Since(start)
	log.Printf("calculated %v with pointer (++) took: %s", value, elapsed)

	value = 5

	start = time.Now()

	for i := 0; i < 1600000; i++ {
		value = plus(value)
	}

	elapsed = time.Since(start)
	log.Printf("calculated %v without pointer took: %s", value, elapsed)
}
