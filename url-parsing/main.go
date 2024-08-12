package main

import (
	"fmt"
	"net/url"
)

func main() {
	url, _ := url.Parse("https://www.google.com/search?q=hello+world")
	fmt.Println(url)
	fmt.Println(url.Path)
	fmt.Println(url.Host)
}
