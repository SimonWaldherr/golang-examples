package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	// it will be 1
	fmt.Println("FOO:", os.Getenv("FOO"))

	// it will be empty string
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	envs := os.Environ()
	fmt.Printf("the number of environment variables: %d\n", len(envs))

	// traverse max 3 env variables
	count := min(3, len(envs))

	for i := range count {
		// print original string
		fmt.Printf("original pair string: %s\n", envs[i])

		// as the environment varaiables format as key=value
		// so we can use [strings.Cut] function
		k, v, found := strings.Cut(envs[i], "=")
		if !found {
			fmt.Printf(" =  not found on kv pair, skip this, str: %s\n", envs[i])
			continue
		}

		fmt.Printf("key: %s, value: %s\n\n", k, v)
	}
}
