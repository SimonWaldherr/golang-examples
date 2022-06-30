package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
