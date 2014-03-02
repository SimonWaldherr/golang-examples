package main

import (
	"fmt"
	"os"
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
	if len(os.Args) == 1 {
		fmt.Println("start this application with the argument true to compute primenumbers parallel or false for serial")
		fmt.Println("you can configure the maximum processes/threads amount with: \"export GOMAXPROCS=$number\"")
	} else {
		if os.Args[1] == "true" {
			prime0 := make(chan int)
			prime1 := make(chan int)
			prime2 := make(chan int)
			prime3 := make(chan int)
			prime4 := make(chan int)
			prime5 := make(chan int)
			prime6 := make(chan int)
			prime7 := make(chan int)
			go func() {
				prime0 <- getPrime(200000)
			}()
			go func() {
				prime1 <- getPrime(500000)
			}()
			go func() {
				prime2 <- getPrime(100000)
			}()
			go func() {
				prime3 <- getPrime(250000)
			}()
			go func() {
				prime4 <- getPrime(550000)
			}()
			go func() {
				prime5 <- getPrime(150000)
			}()
			go func() {
				prime6 <- getPrime(350000)
			}()
			go func() {
				prime7 <- getPrime(300000)
			}()

			for i := 0; i < 8; i++ {
				select {
				case msg0 := <-prime0:
					fmt.Print("the 200000th prime number is: ")
					fmt.Println(msg0)
				case msg1 := <-prime1:
					fmt.Print("the 500000th prime number is: ")
					fmt.Println(msg1)
				case msg2 := <-prime2:
					fmt.Print("the 100000th prime number is: ")
					fmt.Println(msg2)
				case msg3 := <-prime3:
					fmt.Print("the 250000th prime number is: ")
					fmt.Println(msg3)
				case msg4 := <-prime4:
					fmt.Print("the 550000th prime number is: ")
					fmt.Println(msg4)
				case msg5 := <-prime5:
					fmt.Print("the 150000th prime number is: ")
					fmt.Println(msg5)
				case msg6 := <-prime6:
					fmt.Print("the 350000th prime number is: ")
					fmt.Println(msg6)
				case msg7 := <-prime7:
					fmt.Print("the 300000th prime number is: ")
					fmt.Println(msg7)
				}
			}
		}
		if os.Args[1] == "false" {
			fmt.Print("the 200000th prime number is: ")
			fmt.Println(getPrime(200000))
			fmt.Print("the 500000th prime number is: ")
			fmt.Println(getPrime(500000))
			fmt.Print("the 100000th prime number is: ")
			fmt.Println(getPrime(100000))
			fmt.Print("the 250000th prime number is: ")
			fmt.Println(getPrime(250000))
			fmt.Print("the 550000th prime number is: ")
			fmt.Println(getPrime(550000))
			fmt.Print("the 150000th prime number is: ")
			fmt.Println(getPrime(150000))
			fmt.Print("the 350000th prime number is: ")
			fmt.Println(getPrime(350000))
			fmt.Print("the 300000th prime number is: ")
			fmt.Println(getPrime(300000))
		}
	}
}
