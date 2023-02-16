package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	service := ":8090"
	cpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	// TCP サーバーダイアル
	conn, err := net.DialTCP("tcp", nil, cpAddr)
	checkError(err)
	// 送信する
	_, err = conn.Write([]byte("hello!"))
	checkError(err)
	result, err := io.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
