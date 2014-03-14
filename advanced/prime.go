package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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
	flag.Parse()
	s := flag.Arg(0)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("please provide a number")
		os.Exit(2)
	}
	fmt.Println(getPrime(i))
}
