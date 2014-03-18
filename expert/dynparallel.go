package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

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

func main() {
	nprocs, _ := strconv.Atoi(os.Args[1])
	fmt.Println("setting max processes to: ", nprocs)
	runtime.GOMAXPROCS(nprocs)

	var prime int
	channels := make([]chan int, nprocs)
	for i := 0; i < nprocs; i++ {
		channels[i] = make(chan int)
	}

	for i := 0; i < nprocs; i++ {
		go func(ch chan int, val int) {
			fmt.Printf("sending %v\n", val)
			rand.Seed(time.Now().UnixNano())
			prime = getPrime(val + rand.Intn(32))
			ch <- prime
		}(channels[i], i)
	}

	var chancount int = 0
	for {
		for i := 0; i < nprocs; i++ {
			select {
			case v := <-channels[i]:
				fmt.Printf("received %v from channel %v\n", v, i)
				chancount++
			default:
			}
		}
		if chancount == nprocs {
			break
		}

	}
}
