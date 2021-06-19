package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"simonwaldherr.de/go/golibs/as"
)

func defaultValue(va ...interface{}) (int64, int64, int64, int64) {
	var v [4]int64
	var typeof = ""
	for i := 0; i < 4; i++ {
		if i < len(va) {
			typeof = as.String(reflect.TypeOf(va[i]))
		} else {
			typeof = "nil"
		}

		if typeof == "int" || typeof == "int8" || typeof == "int16" || typeof == "int32" || typeof == "int64" {
			v[i] = as.Int(va[i])
		} else {
			switch i {
			case 0:
				v[i] = 2
			case 1:
				v[i] = 4
			case 2:
				v[i] = 8
			case 3:
				v[i] = 16
			}
		}
	}
	return v[0], v[1], v[2], v[3]
}

func askFor(question string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("# %v: ", question)
	scanner.Scan()
	return string(scanner.Text())
}

func askOnRuntime(va ...interface{}) (int64, int64, int64, int64) {
	var v [4]int64
	var typeof = ""
	for i := 0; i < 4; i++ {
		if i < len(va) {
			typeof = as.String(reflect.TypeOf(va[i]))
		} else {
			typeof = "nil"
		}

		if typeof == "int" || typeof == "int8" || typeof == "int16" || typeof == "int32" || typeof == "int64" {
			v[i] = as.Int(va[i])
		} else {
			switch i {
			case 0:
				v[i] = as.Int(askFor("first value"))
			case 1:
				v[i] = as.Int(askFor("second value"))
			case 2:
				v[i] = as.Int(askFor("third value"))
			case 3:
				v[i] = as.Int(askFor("fourth value"))
			}
		}
	}
	return v[0], v[1], v[2], v[3]
}

func optionalValue(v ...interface{}) (int64, error) {
	countArguments := len(v)
	if countArguments < 4 && countArguments > 0 { //we check for the right amount of arguments
		switch countArguments {
		case 1:
			return as.Int(v[0]), nil
		case 2:
			return as.Int(v[0]) * as.Int(v[1]), nil
		case 3:
			return as.Int(v[0]) * as.Int(v[1]) * as.Int(v[2]), nil
		}
	}
	return int64(len(v)), errors.New(fmt.Sprintf("wrong number of arguments, expected between 1 and 3, discovered %v", countArguments))
}

func printAll(v ...interface{}) {
	for _, x := range v {
		fmt.Printf("%v\t", x)
	}
	fmt.Println()
}

func sprintAll(v ...interface{}) string {
	var str = ""
	for _, x := range v {
		str += fmt.Sprintf("%v\t", x)
	}
	str += fmt.Sprintln()
	return str
}

func first(v ...interface{}) interface{} {
	return v[0]
}

func main() {
	fmt.Printf("first(optionalValue(12, 23, 42))\n%v\n", first(optionalValue(12, 23, 42)))
	fmt.Printf("\noptionalValue()\n%v", sprintAll(optionalValue()))
	fmt.Printf("\noptionalValue(23)\n%v", sprintAll(optionalValue(23)))
	fmt.Printf("\noptionalValue(23, 42)\n%v", sprintAll(optionalValue(23, 42)))
	fmt.Printf("\noptionalValue(12, 23, 42)\n%v", sprintAll(optionalValue(12, 23, 42)))
	fmt.Printf("\noptionalValue(1, 2, 3, 10)\n%v", sprintAll(optionalValue(1, 2, 3, 10)))
	fmt.Println("\ndefaultValue(3.1415, 4, 5)")
	printAll(defaultValue(3.1415, 4, 5))
	fmt.Println("\ndefaultValue(\"1\", 3, 5, 7, 9)")
	printAll(defaultValue("1", 3, 5, 7, 9))
	fmt.Println("\ndefaultValue(\"1\", false, nil, 7, 9)")
	printAll(askOnRuntime("1", false, nil, 7, 9))
}
