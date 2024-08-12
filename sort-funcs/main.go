package main

import (
	"log"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.SortFunc(strs, func(a, b string) int {
		if a < b {
			return -1
		}

		if a > b {
			return 1
		}

		return 0
	})
	log.Println(strs)
}
