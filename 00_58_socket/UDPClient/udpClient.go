package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	port := ":9090"

	udpAddr, err := net.ResolveUDPAddr("udp4", port)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	_, err = conn.Write([]byte("anything"))
	checkError(err)

	// res, err := io.ReadAll(conn)
	// checkError(err)
	//
	// fmt.Println(string(res))
	// os.Exit(0)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
