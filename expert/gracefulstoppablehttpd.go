package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

type handler struct {
	quitChan chan os.Signal
	counter  *atomic.Int32
}

func (h *handler) stop(w http.ResponseWriter, r *http.Request) {
	h.quitChan <- os.Interrupt
	w.Write([]byte("listener stopped"))
}

func (h *handler) normal(w http.ResponseWriter, r *http.Request) {
	number := h.counter.Add(1)
	fmt.Fprintf(w, "Hello World\nYou requested: %s\n", r.URL)
	time.Sleep(4 * time.Second)
	fmt.Fprintf(w, "this is request number %d\n", number)
	log.Println(r.URL)
}

func handleQuit(quitChan chan os.Signal, srv *http.Server) {
	<-quitChan

	exitCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(exitCtx)
}

func main() {
	log.Printf("started program")

	srv := &http.Server{
		Addr: ":8080",
	}

	h := &handler{
		quitChan: make(chan os.Signal, 1),
		counter:  new(atomic.Int32),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.normal)
	mux.HandleFunc("/stop", h.stop)
	srv.Handler = mux

	signal.Notify(h.quitChan, os.Interrupt)
	go handleQuit(h.quitChan, srv)

	err := srv.ListenAndServe()
	switch {
	case errors.Is(err, http.ErrServerClosed):
		log.Println("server closed. bye o/")
	case err != nil:
		log.Fatalf("ListenAndServe: %s", err)
	}
}
