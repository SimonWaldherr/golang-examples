package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Home Page!")
	fmt.Println("Endpoint Hit: homePage")
}

func about(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "This repository contains Go(lang) examples.")
	fmt.Println("Endpoint Hit: about")
}

func request() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", about)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	request()
}
