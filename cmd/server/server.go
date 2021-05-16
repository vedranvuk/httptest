package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	listenPort = flag.Int("listenPort", 8080, "Localhost port to listen on")
)

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
	})
	var server = http.Server{
		Addr: fmt.Sprintf("localhost:%d", *listenPort),
		Handler: mux,
	}
	var err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
