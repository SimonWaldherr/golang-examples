// Description: Error handling in Go - creating, returning, and wrapping errors
// Tags: error, errors, fmt.Errorf, error handling, sentinel error
package main

import (
	"errors"
	"fmt"
)

// ErrNegativeNumber is a sentinel error used to signal a specific condition.
var ErrNegativeNumber = errors.New("negative number not allowed")

// sqrt returns the square root of n using Newton's method, or an error.
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("sqrt(%g): %w", n, ErrNegativeNumber)
	}
	// Newton-Raphson approximation
	z := n
	for i := 0; i < 10; i++ {
		z -= (z*z - n) / (2 * z)
	}
	return z, nil
}

// divide returns a / b or an error when b is zero.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	// Successful call
	if result, err := sqrt(16); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("sqrt(16) = %.4f\n", result)
	}

	// Call that returns an error
	if _, err := sqrt(-4); err != nil {
		fmt.Println("error:", err)

		// errors.Is checks the error chain for a sentinel value
		if errors.Is(err, ErrNegativeNumber) {
			fmt.Println("hint: provide a non-negative number")
		}
	}

	// Another error example
	if result, err := divide(10, 3); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	if _, err := divide(5, 0); err != nil {
		fmt.Println("error:", err)
	}
}
