package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var regex string
	if len(os.Args) > 1 {
		regex = os.Args[1]
	} else {
		fmt.Fprintln(os.Stderr, "pipe.go needs at least one parameter")
		os.Exit(1)
	}
	pipe := bufio.NewScanner(os.Stdin)

	for pipe.Scan() {
		str := pipe.Text()
		matched, _ := regexp.MatchString(regex, str)
		if matched {
			fmt.Println(str)
		}
	}

	if err := pipe.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
