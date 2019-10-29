package main

import "fmt"

func main() {
	//simple constants
	const stringConstant = "string"
	const boolConstant = true
	const intConstant = 1234

	fmt.Print(stringConstant)
	fmt.Print(boolConstant)
	fmt.Print(intConstant)

	//simple constant with type
	const float64Constant float64 = 1234.00

	fmt.Print(float64Constant)

	//multiple constants
	const color, code = "red", 255

	fmt.Print(color)
	fmt.Print(code)

	const (
		company string  = "Go Experts"
		salary  float64 = 50000.0
	)
	
	fmt.Print(company)
	fmt.Print(salary)
}
