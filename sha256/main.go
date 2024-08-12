package main

import (
	"crypto/sha256"
	"fmt"
)

// sha256 hashes are frequent used to compute short identity for binary or texts blobs
// For example, TLS/SSL certificates use SHA256 to compute a certificate’s signature

func main() {
	s := "this is string"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
