package main

import (
	"fmt"
)

type Day int

const (
	SUNDAY Day = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

func day() {
	fmt.Printf("Monday has the value %d\n", MONDAY)
	fmt.Printf("Friday has the value %d\n", FRIDAY)
}

type Days int

const (
	SUNDAYS Days = 1 << iota
	MONDAYS
	TUESDAYS
	WEDNESDAYS
	THURSDAYS
	FRIDAYS
	SATURDAYS
)

func days(d Days) {
	fmt.Print("\nYou selected these days:\n")

	if d&SUNDAYS != 0 {
		fmt.Print("* SUNDAY\n")
	}
	if d&MONDAYS != 0 {
		fmt.Print("* MONDAY\n")
	}
	if d&TUESDAYS != 0 {
		fmt.Print("* TUESDAY\n")
	}
	if d&WEDNESDAYS != 0 {
		fmt.Print("* WEDNESDAY\n")
	}
	if d&THURSDAYS != 0 {
		fmt.Print("* THURSDAY\n")
	}
	if d&FRIDAYS != 0 {
		fmt.Print("* FRIDAY\n")
	}
	if d&SATURDAYS != 0 {
		fmt.Print("* SATURDAY\n")
	}
}

func main() {
	day()

	days(34)

	days(MONDAYS | FRIDAYS)

	days(5)

	days(127)
}
