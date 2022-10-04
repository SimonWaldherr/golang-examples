package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "simpleForm.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Printf("Post form arrived !  %v\n", r.PostForm)
		nickname := r.FormValue("nickname")
		password := r.FormValue("password")
		fmt.Printf("Nickname = %s\n", nickname)
		fmt.Printf("Password = %s\n", password)
	default:
		fmt.Fprintf(w, "Sorry! Only GET and POST are suppoterd!")

	}

}

func main() {
	http.HandleFunc("/", handleRequests)

	fmt.Println("Starting server at 8080 !")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
