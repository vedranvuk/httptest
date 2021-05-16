package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	wg          = sync.WaitGroup{}
	serverPort  = flag.Int("serverPort", 8080, "Port of the HTTP server on localhost to connect to (8080).")
	numThreads  = flag.Int("numThreads", 100, "Number of request threads (100).")
	numRequests = flag.Int("numRequests", 1000, "Number of requests to perform per thread (1000).")
)

func requestThread() {
	var client = http.Client{
		Transport: &http.Transport{},
	}
	var resp *http.Response
	var err error
	for i := 0; i < *numRequests; i++ {
		resp, err = client.Get(fmt.Sprintf("http://localhost:%d", *serverPort))
		if err != nil {
			fmt.Println(err)
		}
		resp.Body.Close()
	}
	wg.Done()
}

func main() {
	wg.Add(*numThreads)
	var start = time.Now()
	for i := 0; i < *numThreads; i++ {
		go requestThread()
	}
	wg.Wait()
	fmt.Printf("Operation took %dms.", time.Since(start).Milliseconds())
	http.Get(fmt.Sprintf("http://localhost:%d/stop", *serverPort))
}
