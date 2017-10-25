package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func main() {
	response, err := http.Get("http://golang.org/") // Send a HTTP GET request to the URL
	if err != nil {
		fmt.Printf("%s", err) // If there is an error, print it and exit
		os.Exit(1)
	} else {
		defer response.Body.Close()                    // Close the connection to prevent a resource leak
		contents, err := ioutil.ReadAll(response.Body) // Read the contents of the HTTP GET
		if err != nil {                                // If there is an error
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents)) // Print the contents of the request
	}
}
