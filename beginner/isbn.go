package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	isbnInput := os.Args[1]
	isbnClean := strings.Replace(isbnInput, "-", "", -1)
	fmt.Printf("ISBN verification of %v: %v\n", isbnInput, VerifyISBN(isbnClean))
}

func isNotNumber(n rune) bool {
	return n < '0' || '9' < n
}

func VerifyISBN(code string) bool {
	if len(code) == 10 {
		return VerifyISBN10(code)
	} else if len(code) == 13 {
		return VerifyISBN13(code)
	} else {
		return false
	}
}

func VerifyISBN10(code string) bool {

	if len(code) != 10 {
		return false
	}

	sum, multiply := 0, 10
	for _, n := range code {

		var d int
		switch {
		case n == 'X':
			d = 10
		case isNotNumber(n):
			return false
		default:
			d = int(n - '0')
		}

		sum += multiply * d
		multiply--
	}

	return sum%11 == 0
}

func VerifyISBN13(code string) bool {

	if len(code) != 13 {
		return false
	}

	sum, weight := 0, 1
	for _, n := range code[:len(code)-1] {
		if isNotNumber(n) {
			return false
		}
		sum += int(n-'0') * weight
		if weight == 1 {
			weight = 3
		} else {
			weight = 1
		}
	}

	d := 10 - sum%10
	if d == 10 {
		d = 0
	}

	return d == int(code[len(code)-1]-'0')
}
