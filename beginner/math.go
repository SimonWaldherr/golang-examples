package main

import (
	"fmt"
)

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func GreatestCommonDivisor(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func LeastCommonMultiple(a, b int) int {
	return a * b / GreatestCommonDivisor(a, b)
}

func Sqrt(n int64) int64 {
	var t int64
	var b int64
	var r int64
	t = int64(n)
	p := int64(1 << 30)
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
	return int64(r)
}

func Prime(n int) int {
	var primeList = []int{2}
	isPrime := 1
	num := 3
	sqrtNum := 0
	for len(primeList) < n {
		sqrtNum = int(Sqrt(int64(num)))
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

func Sum(numbers ...int) int {
	var sum int
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func main() {
	fmt.Printf("Abs(-3): %v\n", Abs(-3))
	fmt.Printf("GreatestCommonDivisor(6,16): %v\n", GreatestCommonDivisor(6, 16))
	fmt.Printf("LeastCommonMultiple(12,24): %v\n", LeastCommonMultiple(12, 24))
	fmt.Printf("Sqrt(9): %v\n", Sqrt(9))
	fmt.Printf("Prime(15): %v\n", Prime(15))
	fmt.Printf("Sum(3,5,7,9): %v\n", Sum(3, 5, 7, 9))
}
