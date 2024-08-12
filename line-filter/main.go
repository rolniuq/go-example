package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// COMMANDS:
// echo 'hello' > /tmp/lines
// echo 'filter' >> /tmp/lines
//  cat /tmp/lines | go run main.go

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("error: ", err)
		os.Exit(1)
	}
}
