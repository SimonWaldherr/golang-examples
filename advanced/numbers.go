package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var input string
	var lstc1 string
	var lstc2 int64

	if len(os.Args) >= 2 {
		input = os.Args[1]
	} else {
		fmt.Println("Enter a number")
		_, err1 := fmt.Scanf("%v", &input)
		check(err1)
	}

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
		fmt.Printf("%vth\n", intin)
		break
	case lstc1 == "1":
		fmt.Printf("%vst\n", intin)
		break
	case lstc1 == "2":
		fmt.Printf("%vnd\n", intin)
		break
	case lstc1 == "3":
		fmt.Printf("%vrd\n", intin)
		break
	default:
		fmt.Printf("%vth\n", intin)
	}
}
