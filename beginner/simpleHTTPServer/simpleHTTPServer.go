package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Log is a middleware print the RemoteAddr, Method and URL before handle this request
// You should use it like a decorator
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	// default settings
	port := "8080"
	// dir := os.Getenv("PWD")
	// FIX: Use [os.Getwd] to ensure cross-platform valid
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("get current dir failed, using . as serve dir, err: %s", err)
		dir = "."
	}

	// get settings from command line
	// the first additional argument is port, the second is serve dir
	// by default, the port is 8080 and the serve dir is current dir
	//
	// NOTE: This doesn't parse flag arguments like --port -p
	if len(os.Args) > 1 {
		port = os.Args[1]
		if len(os.Args) > 2 {
			dir = os.Args[2]
		}
	}

	// start web server with logging
	fmt.Printf("Serving HTTP on 0.0.0.0 port %v ...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, Log(http.FileServer(http.Dir(dir)))))
}
