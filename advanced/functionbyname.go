package main

import (
	"fmt"
)

func f1() string {
	return "foo"
}
func f2() string {
	return "bar: "
}
func f3() string {
	return "lorem "
}
func f4() string {
	return "ipsum "
}
func f5() string {
	return "dolor "
}
func f6() string {
	return "sit "
}
func f7() string {
	return "amet."
}

func main() {
	funcs := map[string]func() string{
		"f1": f1,
		"f2": f2,
		"f3": f3,
		"f4": f4,
		"f5": f5,
		"f6": f6,
		"f7": f7,
	}

	var str string
	for i := 1; i < 8; i++ {
		x := fmt.Sprintf("%v", i)
		str += funcs["f"+x]()
	}
	fmt.Println(str)
}
