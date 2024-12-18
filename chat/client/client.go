package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("Error", err)
				break
			}

			fmt.Println("Received Msg", string(msg))

		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			if err := c.WriteMessage(websocket.TextMessage, []byte("PING")); err != nil {
				log.Println("Error", err)
				break
			}

			fmt.Println("Wrote PING Msg")
		}
	}
}
