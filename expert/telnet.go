package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(c net.Conn, msgchan chan<- string) {
	buf := make([]byte, 4096)
	for {
		n, err := c.Read(buf)
		if (err != nil) || (n == 0) {
			c.Close()
			break
		}
		msgchan <- string(buf[0:n])
	}
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}

func printMessages(msgchan <-chan string) {
	for {
		msg := strings.TrimSpace(<-msgchan)
		fmt.Printf("data: %s\n", msg)
	}
}

func main() {
	flag.Parse()
	port := ":" + flag.Arg(0)
	if port == ":" {
		port = ":23"
	}
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	msgchan := make(chan string)
	go printMessages(msgchan)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn, msgchan)
	}
}
