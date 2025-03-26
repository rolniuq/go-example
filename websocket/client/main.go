package client

import (
	"context"
	"fmt"
	"net/http"
)

type WsClient struct {
	ctx  context.Context
	port int64
}

func NewWsClient(ctx context.Context, port int64) *WsClient {
	return &WsClient{
		port: port,
	}
}

func (ws *WsClient) openServer() error {
	http.HandleFunc("/ws-client", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from client"))
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", ws.port), nil); err != nil {
		return err
	}

	return nil
}

func (ws *WsClient) Start() error {
	return ws.openServer()
}

func (ws *WsClient) Stop() {

}
