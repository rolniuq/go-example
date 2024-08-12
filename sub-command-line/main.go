package main

import (
	"flag"
	"fmt"
	"os"
)

// like go tool or git has their subcommands
//  ex: go tool have go get, go build, ...
// flag package will help to define subcommand

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected foo or bad subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand foo")
		fmt.Println("enable", *fooEnable)
		fmt.Println("name", *fooName)
		fmt.Println("tail", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand bar")
		fmt.Println("level", *barLevel)
		fmt.Println("tail", barCmd.Args())
	default:
	}
}

// go build main.go
// ./main foo -enable -name=joe a1 a2
// ./main bar -level 8 a1
