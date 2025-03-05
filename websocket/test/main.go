package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Echo server: Reads a message and sends it back
func wsHandler(ws *websocket.Conn) {
	defer ws.Close()

	for {
		var msg string
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Println("WebSocket closed:", err)
			break
		}

		fmt.Println("Received:", msg)

		// Send the same message back to client
		err = websocket.Message.Send(ws, "Echo: "+msg)
		if err != nil {
			log.Println("Send error:", err)
			break
		}
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(wsHandler))

	port := "8080"
	fmt.Println("WebSocket server listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
