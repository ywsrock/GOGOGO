package main

import (
	"context"
	"fmt"
	"log"

	v "myBuf/gen/proto/go/dic/v1"

	"google.golang.org/grpc"
)

func main() {
	connectTo := "0.0.0.0:9090"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := v.NewDicServiceClient(conn)
	req := v.GetKeyRequest{
		Name: "Client",
	}

	res, err := c.GetKey(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
