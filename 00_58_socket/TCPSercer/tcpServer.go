package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	port := ":8090"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	chkError(err)
	// 受信ポート監視
	listener, err := net.ListenTCP("tcp", tcpAddr)
	chkError(err)
	fmt.Println("TCP start Success!")

	for {
		// 接続待ち
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleRequest(conn)

	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	request := make([]byte, 255)
	for {
		// 受信情報
		l, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if l == 0 {
			fmt.Println(l)
			break
		} else {
			w := bufio.NewWriter(conn)
			w.WriteString(time.Now().Format("2006-01-02"))
			w.Flush()

		}
		conn.Close()
		request = make([]byte, 255)
	}
}
func chkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
