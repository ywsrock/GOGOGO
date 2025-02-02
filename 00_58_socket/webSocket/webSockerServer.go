package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", websocket.Handler(ScHandl))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func ScHandl(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)
		go func() {
			for {
				select {
				case _ = <-time.Tick(5 * time.Second):
					if err = websocket.Message.Send(ws, msg); err != nil {
						fmt.Println("Can't send")
						break
					}
				}
			}
		}()

	}
}
