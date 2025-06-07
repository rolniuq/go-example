package server

import (
	"context"
	"fmt"
	"go-practice/day14/greeter"
	"go-practice/day14/interceptor"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct{}

type server struct {
	greeter.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	traceID := "unknown"
	if ok {
		values := md["trace-id"]
		if len(values) > 0 {
			traceID = values[0]
		}
	}

	fmt.Println("Trace ID from client:", traceID)

	return &greeter.HelloReply{
		Message: fmt.Sprintf("[trace-id %s] Hello %s", traceID, req.Name),
	}, nil
}

func (s *Server) Exec() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	grpc := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.TraceServerInterceptor()),
	)
	greeter.RegisterGreeterServer(grpc, &server{})

	fmt.Println("start server on port 50051")
	if err := grpc.Serve(lis); err != nil {
		panic(err)
	}
}
