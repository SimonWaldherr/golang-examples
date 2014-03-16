package main

import (
	"os"
	"fmt"
	"strconv"
	"os/exec"
)

func main() {
	pid := os.Getpid()
	str := strconv.Itoa(pid)
	fmt.Println("Process identifier: ", str)
	ret, _ := exec.Command("kill", "-9", str).Output()
	fmt.Println("this will never be printed: ", ret)
}
