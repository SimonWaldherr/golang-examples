package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var listener net.Listener
var requestcounter int
var stopvar bool

type handler struct {
	wg sync.WaitGroup
	*http.ServeMux
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.wg.Add(1)
	defer h.wg.Done()
	h.ServeMux.ServeHTTP(w, r)

	w.(http.Flusher).Flush()
}

func restart(w http.ResponseWriter, r *http.Request) {
	listener.Close()
	w.Write([]byte("listener restarted"))
}

func stop(w http.ResponseWriter, r *http.Request) {
	stopvar = true
	listener.Close()
	w.Write([]byte("listener stoped"))
}

func normal(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if r.URL.RawQuery != "" {
		url = strings.Join([]string{url, "?", r.URL.RawQuery}, "")
	}
	fmt.Fprintf(w, "Hello World\nYou requested: %s\n", url)
	time.Sleep(4 * time.Second)
	fmt.Fprintf(w, "this is request number %s\n", strconv.Itoa(requestcounter))
	log.Println(url)
	requestcounter++
}

func main() {
	log.Printf("started program")
	var err error
	var ler string
	for !stopvar {
		listener, err = net.Listen("tcp", ":8080")

		if err == nil {
			log.Printf("started listener on port 8080 with counter: %v\n", requestcounter)
			h := &handler{ServeMux: http.NewServeMux()}

			h.ServeMux.HandleFunc("/", normal)
			h.ServeMux.HandleFunc("/stop", stop)
			h.ServeMux.HandleFunc("/restart", restart)

			http.Serve(listener, h)

			h.wg.Wait()
		} else {
			if fmt.Sprintf("%v", err) != ler {
				ler = fmt.Sprintf("%v", err)
				fmt.Println(ler)
			}
		}
	}
}
