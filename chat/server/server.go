package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		defer c.Close()

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("Error", err)
				break
			}

			fmt.Println("Received Msg", string(msg))

			if string(msg) == "PING" {
				if err := c.WriteMessage(websocket.TextMessage, []byte("PONG")); err != nil {
					log.Println("Error", err)
					break
				}

				fmt.Println("Wrote Msg PONG")
			}
		}
	})

	fmt.Println("Server is listening on port 3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
