package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func handleConnection(c net.Conn, msgchan chan<- string) {
	defer c.Close()
	fmt.Printf("Connection from %v established.\n", c.RemoteAddr())
	c.SetReadDeadline(time.Now().Add(time.Second * 5))
	buf := make([]byte, 4096)
	for {
		n, err := c.Read(buf)
		if (err != nil) || (n == 0) {
			c.Close()
			break
		}
		msgchan <- string(buf[0:n])
	}
	time.Sleep(150 * time.Millisecond)
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
	c.Close()
	return
}

func printMessages(msgchan <-chan string) {
	var count int = 0
	for {
		msg := strings.TrimSpace(<-msgchan)
		count++
		fmt.Printf("Data %d: %s\n", count, msg)
	}
}

func main() {
	flag.Parse()
	port := ":" + flag.Arg(0)
	if port == ":" {
		port = ":2223"
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
