package main

import (
	"fmt"
	"os"
)

var (
	// list of nth prime numbers to find
	requests = []int{200000, 500000, 100000, 250000, 550000, 150000, 350000, 300000}
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("start this application with the argument true to compute primenumbers parallel or false for serial")
		fmt.Println("you can configure the maximum processes/threads amount with: \"export GOMAXPROCS=$number\"")
		os.Exit(1)
	}

	if os.Args[1] == "true" {
		runParallel()
	} else {
		runSequential()
	}
}

func runSequential() {
	for _, index := range requests {
		fmt.Printf("the %dth prime number is: %d\n", index, getPrime(index))
	}
}

func runParallel() {
	// data struct that goroutines will send information
	// back to main thread
	type WorkerResponse struct {
		Index int
		Prime int
	}

	workerChan := make(chan WorkerResponse)
	defer close(workerChan)

	// send requests to n goroutines
	for _, index := range requests {
		// start this goroutine with the index in the loop
		// we must give this param, because index would be shared memory
		go func(idx int) {
			workerChan <- WorkerResponse{Index: idx, Prime: getPrime(idx)}
		}(index)
	}

	for i := 0; i < len(requests); i++ {
		response := <-workerChan
		fmt.Printf("the %dth prime number is: %d\n", response.Index, response.Prime)
	}
}

func Sqrt(n int) int {
	var t uint
	var b uint
	var r uint
	t = uint(n)
	p := uint(1 << 30)
	for p > t {
		p >>= 2
	}
	for ; p != 0; p >>= 2 {
		b = r | p
		r >>= 1
		if t >= b {
			t -= b
			r |= p
		}
	}
	return int(r)
}

func getPrime(n int) int {
	var primeList = []int{2}
	var isPrime int = 1
	var num int = 3
	var sqrtNum int = 0
	for len(primeList) < n {
		sqrtNum = Sqrt(num)
		for i := 0; i < len(primeList); i++ {
			if num%primeList[i] == 0 {
				isPrime = 0
			}
			if primeList[i] > sqrtNum {
				i = len(primeList)
			}
		}
		if isPrime == 1 {
			primeList = append(primeList, num)
		} else {
			isPrime = 1
		}
		num = num + 2
	}
	return primeList[n-1]
}
