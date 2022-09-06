package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello2 "hello2/proto"
	"log"
	"time"
)

func main()  {

	conn,err := grpc.Dial(":8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	c := hello2.NewHello2Client(conn)

	ctx,cancel := context.WithTimeout(context.Background(),time.Second)

	defer cancel()

	r, err := c.SayName(ctx,&hello2.MsgReq{Msg: "yyyyyy"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(r.GetMsg())

}
