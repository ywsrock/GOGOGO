package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello2 "hello2/proto"
	"log"
	"net"

)

type Server struct{
	hello2.UnimplementedHello2Server
}

func (s *Server) SayName(ctx context.Context, req *hello2.MsgReq) (*hello2.MsgRsp, error) {
	fmt.Println(req.GetMsg())
	res := &hello2.MsgRsp{Msg: "SSSSSS"}

	return res,nil
}

func main()  {
	l,err := net.Listen("tcp",":8080")

	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	hello2.RegisterHello2Server(s,new(Server))

	if err := s.Serve(l);err != nil {
		log.Fatalln("起動失敗")
	}
	log.Println("サーバー起動")

}