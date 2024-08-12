package main

import (
	"flag"
	"fmt"
)

// used: flag.Int("w"), flag.String("s")
// with run command: -w=1 -s=hello

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("folk", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string bar")

	flag.Parse()

	fmt.Println("word", *wordPtr)
	fmt.Println("num", *numPtr)
	fmt.Println("folk", *forkPtr)
	fmt.Println("svar", svar)
	fmt.Println("tail", flag.Args())
}
