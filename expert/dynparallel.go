package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
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

func enNmbr(input string) string {
	var lstc1 string
	var lstc2 int64
	intin, err2 := strconv.ParseInt(input, 10, 0)
	check(err2)

	lstc1 = input[len(input)-1:]
	if intin >= 10 {
		lstc2, _ = strconv.ParseInt(input[len(input)-2:], 10, 0)
	} else {
		lstc2 = 0
	}
	lstc1 = input[len(input)-1:]

	switch {
	case lstc2 > 10 && lstc2 < 20:
		return input + "th"
	case lstc1 == "1":
		return input + "st"
	case lstc1 == "2":
		return input + "nd"
	case lstc1 == "3":
		return input + "rd"
	default:
		return input + "th"
	}
	return ""
}

func main() {
	var nprocs int
	if len(os.Args) >= 2 {
		nprocs, _ = strconv.Atoi(os.Args[1])
	} else {
		fmt.Println("Enter a number")
		_, err1 := fmt.Scanf("%d", &nprocs)
		check(err1)
	}
	if nprocs < 2 {
		panic("input to low")
	}

	fmt.Println("setting max processes to: ", nprocs)
	runtime.GOMAXPROCS(nprocs)

	var prime int
	var random int
	channels := make([]chan int, nprocs)
	for i := 0; i < nprocs; i++ {
		channels[i] = make(chan int)
	}

	for i := 0; i < nprocs; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		random = rand.Intn(6400) + 1
		fmt.Printf("sending request for the %v prime number on channel %v\n", enNmbr(strconv.Itoa(random)), i)
		go func(ch chan int, random int, i int) {
			time.Sleep(time.Duration(random) * time.Millisecond)
			fmt.Printf("calculating the %v prime number for channel %v\n", enNmbr(strconv.Itoa(random)), i)
			prime = getPrime(random)
			ch <- prime
		}(channels[i], random, i)
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
