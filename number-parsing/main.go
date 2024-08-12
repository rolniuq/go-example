package main

import (
	"fmt"
	"strconv"
)

func main() {
	res, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(res)
}
