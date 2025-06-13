package client

import (
	"context"
	"fmt"
	"go-practice/day14/greeter"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Client struct{}

func (c *Client) Exec() {
	conn, err := grpc.NewClient("localhost:50051")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := greeter.NewGreeterClient(conn)

	traceID := uuid.NewString()
	md := metadata.Pairs("trace-id", traceID)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r, err := client.SayHello(ctx, &greeter.HelloRequest{
		Name: "quynh",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Greeting Msg: %s", r.Message)
}
