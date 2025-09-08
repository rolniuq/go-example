package main

import (
	"bufio"
	"fmt"
	"httpfromtcp/config"
	"log"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", config.PORT))
	if err != nil {
		log.Fatal("failed to resolve udp addr", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal("failed to dial udp", err.Error())
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error reading string:", err.Error())
			os.Exit(1)
		}

		if _, err := conn.Write([]byte(message)); err != nil {
			log.Fatal("error write message", err.Error())
			os.Exit(1)
		}

		fmt.Println(message)
	}
}
