package main

import (
    "log"
    "net/http"
    "os"
)

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
        handler.ServeHTTP(w, r)
    })
}


func main() {
    // default settings
    port := "8080"
    dir := os.Getenv("PWD")

    // get settings from command line
    if len(os.Args) > 1 {
        port = os.Args[1]
        if len(os.Args) > 2 {
            dir = os.Args[2]
        }
    }

    // start web server with logging
    log.Fatal(http.ListenAndServe(":"+port, Log(http.FileServer(http.Dir(dir)))))
}
