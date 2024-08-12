package main

import (
	"fmt"
	"net/http"
)

// http server is meaning we are server and client will send request to us

func sayHello(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("hello")
}

func headerFunc(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, header := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, header)
		}
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/headers", headerFunc)

	http.ListenAndServe(":8080", nil)
}
