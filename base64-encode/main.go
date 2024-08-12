package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "this is a string"

	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(encoded)

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))
}
