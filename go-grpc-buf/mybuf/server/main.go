package main

import (
	"context"
	"fmt"
	"log"
	v "myBuf/gen/proto/go/dic/v1"
	"net"

	"google.golang.org/grpc"
)

type DicServerService struct {
	v.UnimplementedDicServiceServer
}

func main() {
	e := Run()
	if e != nil {
		log.Fatal(e)
	}
}

func Run() error {
	listen := "0.0.0.0:9090"
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal("1")
	}
	server := grpc.NewServer()
	v.RegisterDicServiceServer(server, &DicServerService{})
	er := server.Serve(listener)
	if er != nil {
		log.Fatal(err)
	}
	return nil
}

func (s *DicServerService) GetKey(ctx context.Context, req *v.GetKeyRequest) (*v.GetKeyResponse, error) {
	str := req.GetName()
	res := v.GetKeyResponse{}
	res.Value = fmt.Sprintf("%s%s", "--------------------------->", str)
	return &res, nil
}
