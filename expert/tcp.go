package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const retrytime = 510 * time.Millisecond

func tcpSend(str, addr string) string {
	return tcpSendHelper(str, addr, 3)
}

func tcpSendHelper(str, addr string, retry int) string {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err == nil {
		defer conn.Close()

		payloadBytes := []byte(fmt.Sprintf("%s\r\n\r\n", str))
		if _, err = conn.Write(payloadBytes); err != nil {
			log.Println(err)

			if retry > 0 {
				time.Sleep(retrytime)
				return tcpSendHelper(str, addr, retry-1)
			}
		}
		message, _ := bufio.NewReader(conn).ReadString('\t')
		return fmt.Sprint(message)
	}
	log.Println(err)

	if retry > 0 {
		time.Sleep(retrytime)
		return tcpSendHelper(str, addr, retry-1)
	}
	return ""
}

func main() {
	var host string = "google.de"
	var port string = "80"
	var request string = "GET /index.html"

	response := tcpSend(request, host+":"+port)
	fmt.Printf("response from %v:%v for the request \"%v\" is: %v\n", host, port, request, response)
}
