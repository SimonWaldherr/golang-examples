package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const retrytime = 510 * time.Millisecond

func tcpClient(str, addr string) string {
	return tcpClientHelper(str, addr, 3)
}

func tcpClientHelper(str, addr string, retry int) string {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err == nil {
		defer conn.Close()

		payloadBytes := []byte(fmt.Sprintf("%s\r\n\r\n", str))
		if _, err = conn.Write(payloadBytes); err != nil {
			log.Println(err)

			if retry > 0 {
				time.Sleep(retrytime)
				return tcpClientHelper(str, addr, retry-1)
			}
		}
		bytes := make([]byte, 65535)
		_, err := bufio.NewReader(conn).Read(bytes)
		if err == nil || err == io.EOF {
			return string(bytes)
		}
		return fmt.Sprint(err)
	}
	log.Println(err)

	if retry > 0 {
		time.Sleep(retrytime)
		return tcpClientHelper(str, addr, retry-1)
	}
	return ""
}

func main() {
	var host string = "google.de"
	var port string = "80"
	var request string = "GET /index.html"

	response := tcpClient(request, host+":"+port)
	fmt.Printf("response from %v:%v for the request \"%v\" is: \n%v\n", host, port, request, response)
}
