package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello "hello/proto"
	"log"
	"net"
)

type Server struct {
	hello.UnimplementedMessageServer
}

func (s *Server) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {

	str := request.Word
	fmt.Println(str)

	r := &hello.Result{Name: "yws", Age: 12}
	return &hello.HelloResponse{Result: r}, nil
}

func main() {
	s := &Server{}
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	hello.RegisterMessageServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
