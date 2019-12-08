package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	grpcusage "helloworld/api/proto"
	"log"
	"net"
)

const (
	port = ":50051"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *grpcusage.HelloRequest) (*grpcusage.HelloReply, error) {
	return &grpcusage.HelloReply{
		Message: "hello " + in.Name,
	}, nil
}

func main() {
	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("grpc server listening at: 50051 port")
	server := grpc.NewServer()
	grpcusage.RegisterHelloServer(server, &Server{})
	server.Serve(conn)
}
