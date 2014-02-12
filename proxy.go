package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var url string

func handler(w http.ResponseWriter, r *http.Request) {
	url = strings.Join([]string{"http://", "blog.fefe.de", r.URL.Path}, "")
	if r.URL.RawQuery != "" {
		url = strings.Join([]string{url, "?", r.URL.RawQuery}, "")
	}
	response, err := http.Get(url)
	if err != nil {
		os.Exit(2)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Fprintf(w, string(contents))
		log.Println(url)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
