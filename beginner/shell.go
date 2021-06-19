package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	out, err := exec.Command("echo", "Hello", "world").Output()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", out)
	}
}
