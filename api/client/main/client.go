package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	grpcusage "helloworld/api/proto"
	"log"
	"os"
)

const (
	address     = "localhost:50051"
	defaultName = "牛利兵"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := grpcusage.NewHelloClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	request, err := client.SayHello(context.Background(), &grpcusage.HelloRequest{Name: name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(request.Message)
}
