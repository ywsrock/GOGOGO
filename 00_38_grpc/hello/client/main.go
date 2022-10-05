package main

import (
	"context"
	"google.golang.org/grpc"
	hello "hello/proto"
	"log"
	"time"
)

func main() {

	conn,err := grpc.Dial("localhost:8080",grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := hello.NewMessageClient(conn)
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	r,err := c.SayHello(ctx,&hello.HelloRequest{Word: "---------TTT"})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("From Server:",r.Result.Name)
}