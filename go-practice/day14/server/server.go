package server

import (
	"fmt"
	"go-practice/day14/greeter"
	"net"

	"google.golang.org/grpc"
)

type Server struct{}

type server struct {
	greeter.UnimplementedGreeterServer
}

func (s *Server) Exec() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpc := grpc.NewServer()
	greeter.RegisterGreeterServer(grpc, &server{})

	fmt.Println("start server on port 50051")
	if err := grpc.Serve(lis); err != nil {
		panic(err)
	}
}
