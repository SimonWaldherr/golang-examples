// Description: Switch statement in Go - type switch, expression switch, fallthrough
// Tags: switch, case, fallthrough, type switch, default
package main

import "fmt"

func dayKind(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "weekend"
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "weekday"
	default:
		return "unknown"
	}
}

func grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func describe(i interface{}) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("integer: %d", v)
	case float64:
		return fmt.Sprintf("float64: %g", v)
	case string:
		return fmt.Sprintf("string: %q", v)
	case bool:
		return fmt.Sprintf("bool: %t", v)
	default:
		return fmt.Sprintf("unknown type: %T", v)
	}
}

func main() {
	// Expression switch
	for _, day := range []string{"Monday", "Saturday", "Holiday"} {
		fmt.Printf("%s is a %s\n", day, dayKind(day))
	}

	// Condition switch (no expression after switch keyword)
	for _, score := range []int{95, 83, 72, 61, 45} {
		fmt.Printf("score %d -> grade %s\n", score, grade(score))
	}

	// Type switch
	values := []interface{}{42, 3.14, "hello", true, []int{1, 2, 3}}
	for _, v := range values {
		fmt.Println(describe(v))
	}

	// Fallthrough: execution continues to the next case unconditionally
	n := 1
	switch n {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two (reached via fallthrough)")
	case 3:
		fmt.Println("three")
	}
}
