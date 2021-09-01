package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"
)

var listenConfig = net.ListenConfig{
	Control: func(network, address string, c syscall.RawConn) error {
		var opErr error
		if err := c.Control(func(fd uintptr) {
			opErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEPORT, 1)
		}); err != nil {
			return err
		}
		return opErr
	},
}

func main() {
	pid := os.Getpid()
	listener, err := listenConfig.Listen(context.Background(), "tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	server := &http.Server{}
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "Hello from PID %d \n", pid)
		fmt.Printf("serving to %v (%v) with PID %d \n", req.RemoteAddr, req.Header.Get("X-Forwarded-For"), pid)
	})
	fmt.Printf("HTTP Server with PID %d is running \n", pid)

	panic(server.Serve(listener))
}
