package main

import (
	"errors"
	"fmt"
	"httpfromtcp/config"
	"io"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer f.Close()
		defer close(ch)

		line := ""
		for {
			b := make([]byte, 8)
			if _, err := f.Read(b); err != nil {
				if line != "" {
					ch <- line
				}

				if errors.Is(err, io.EOF) {
					break
				}
			}

			slices := strings.Split(string(b), "\n")
			line += slices[0]

			if len(slices) == 2 {
				ch <- line
				line = slices[1]
			}
		}
	}()

	return ch
}

func main() {
	net, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", config.PORT))
	if err != nil {
		panic(err)
	}
	defer net.Close()

	fmt.Println("TCP is opening on port:", 42069)

	for {
		conn, err := net.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println("Accepted connection from", conn.RemoteAddr())

		ch := getLinesChannel(conn)
		for c := range ch {
			fmt.Println("read:", c)
		}

		fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}

}
