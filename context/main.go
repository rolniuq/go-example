package main

import (
	"fmt"
	"net/http"
	"time"
)

// http server is useful for demonstrating the usage of context.
// context for controlling cancellation
// a context carries deadlines, cancellation signals, and other request-scoped values accoss api boundaries and goroutines

func sayHello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("say hello controller started")
	defer fmt.Println("say hello controller enedde")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":8080", nil)
}

// go run main.go & curl localhost:8080/hello
