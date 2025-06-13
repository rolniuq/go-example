package client

import (
	"context"
	"fmt"
	"go-practice/day14/greeter"
	"go-practice/day14/interceptor"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Client struct{}

func (c *Client) Exec() {
	md := metadata.Pairs("trace-id", uuid.New().String())
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	conn, err := grpc.NewClient(
		"127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.TraceClientInterceptor()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := greeter.NewGreeterClient(conn)
	response, err := client.SayHello(ctx, &greeter.HelloRequest{
		Name: "severus",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Message)
}
