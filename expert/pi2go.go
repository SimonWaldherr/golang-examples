package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"simonwaldherr.de/go/golibs/xmath"
	"syscall"
)

type typeOfCalculation int

const (
	NULL typeOfCalculation = iota
	plus
	minus
	times
	divided
)

func waitForCtrlC() chan os.Signal {
	var signalChannel chan os.Signal
	signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	return signalChannel
}

func multipleOfFour(i int) bool {
	if i%4 == 0 {
		return true
	}
	return false
}

func πByLeibniz(stop chan os.Signal) float64 {
	var calculate = true
	go func() {
		<-stop
		calculate = false
	}()

	var π float64 = 1
	var toc = minus
	var denominator float64 = 3

	for calculate {
		if toc == minus {
			π = π - (1 / denominator)
			denominator += 2
			toc = plus
		} else {
			π = π + (1 / denominator)
			denominator += 2
			toc = minus
		}
	}
	return π * 4
}

func πByEuler(stop chan os.Signal) float64 {
	var calculate = true
	go func() {
		<-stop
		calculate = false
	}()

	var π float64 = 0
	var denominator float64 = 1

	for calculate {
		π = π + 1/(denominator*denominator)
		denominator++
	}
	return math.Sqrt(π * 6)
}

func πByPrime(stop chan os.Signal) float64 {
	var calculate = true
	go func() {
		<-stop
		calculate = false
	}()

	var π float64 = 1
	var denominator int = 2
	var prime int

	for calculate {
		prime = xmath.Prime(denominator)
		if multipleOfFour(prime - 1) {
			π = π * (1 + 1/float64(prime))
		} else {
			π = π * (1 - 1/float64(prime))
		}
		denominator++
	}
	return 2 / π
}

func πByGolang() float64 {
	return math.Pi
}

func main() {
	var algo string = ""
	stopChan := waitForCtrlC()

	if len(os.Args) > 1 {
		algo = os.Args[1]
	}

	switch algo {
	default:
		fmt.Print(πByGolang())
	case "leibniz", "Leibniz":
		fmt.Print(πByLeibniz(stopChan))
	case "euler", "Euler":
		fmt.Print(πByEuler(stopChan))
	case "prime", "Prime":
		fmt.Print(πByPrime(stopChan))
	}

	fmt.Print("\n")
}
