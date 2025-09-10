package nats

import "github.com/nats-io/nats.go"

type Nats struct {
	conn *nats.Conn
}

func NewNats() *Nats {
	conn, err := nats.Connect()
	return &Nats{}
}
