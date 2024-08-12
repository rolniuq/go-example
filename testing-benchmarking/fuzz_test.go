package main

import (
	"testing"
)

// Reverse returns the reversed version of the input string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FuzzReverse(f *testing.F) {
	// Seed inputs
	f.Add("hello")
	f.Add("world")
	f.Add("12345")

	// Fuzz test with generated inputs
	f.Fuzz(func(t *testing.T, s string) {
		reversed := Reverse(s)
		doubleReversed := Reverse(reversed)

		// Check that reversing the string twice gives the original string
		if s != doubleReversed {
			t.Errorf("Original: %q, Double reversed: %q", s, doubleReversed)
		}

		// Optional: Check for specific characteristics, like avoiding empty strings in certain cases
		if len(s) > 0 && reversed == "" {
			t.Errorf("Reverse of %q resulted in an empty string", s)
		}
	})
}
