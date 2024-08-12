package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// support http client and server in the net/http package
// http client meaning we are client and we will send a request to server
// for this example server is gobyexample.com

func main() {
	// send a get request
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
