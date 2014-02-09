package main

import (
  "fmt";
  "net/http";
  "log";
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World\nYou requested: /%s", r.URL.Path[1:]);
  log.Println(r.URL.Path[1:]);
}

func main() {
  http.HandleFunc("/", handler);
  err := http.ListenAndServe(":8080", nil);
  if (err != nil) {
    log.Fatal("ListenAndServe: ", err);
  }
}