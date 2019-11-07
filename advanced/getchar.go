package main

import (
	"fmt"
	"os"
	"os/exec"
)

func readStdin(out chan string, in chan bool) {
	//no buffering
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	//no visible output
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	for {
		select {
		case <-in:
			return
		default:
			os.Stdin.Read(b)
			fmt.Printf(">>> %v: ", b)
			out <- string(b)
		}
	}
}

func main() {
	defer func() {
		exec.Command("stty", "-f", "/dev/tty", "echo").Run()
	}()
	stdin := make(chan string, 1)
	kill := make(chan bool, 1)
	var count int = 0
	var input string

	fmt.Print("this program don't wait for a enter-key\nyou can now start typing\npress \"q\" to quit\n\n")

	go readStdin(stdin, kill)
	for {
		str := <-stdin

		if str == "q" {
			kill <- true
			close(stdin)
			break
		}

		fmt.Println(str)
		input += str
		count++

		if count%10 == 0 {
			fmt.Printf("\n\nyou typed %v keys, quit the demo by pressing \"q\"\n\n", count)
			if count == 600 {
				fmt.Printf("Don't waste your time! Do something meaningful!")
				break
			}
		}
	}
	fmt.Printf("\n\nyou typed %v keys before you quit the program\n", count)
	fmt.Println(input)
}
