package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type tcpServer struct{}

// open begins listening for incoming TCP connections on port 3000.
func (fs *tcpServer) open() {
	// Create a TCP listener on port 3000.
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln(err)
	}

	// Continuously accept new connections and start a new goroutine to handle each connection.
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go fs.handle(conn)
	}
}

// handle reads data from the established connection.
func (fs *tcpServer) handle(conn net.Conn) {
	// Using a buffer to hold incoming data.
	buf := new(bytes.Buffer)
	for {
		// Read the size of the incoming file from the connection.
		var fileSize int64
		err := binary.Read(conn, binary.LittleEndian, &fileSize)
		if err != nil {
			log.Fatal(err)
		}

		// Read the incoming file data into the buffer.
		n, err := io.CopyN(buf, conn, fileSize)
		if err != nil {
			log.Fatal(err)
		}

		// Print the number of bytes received and the actual bytes.
		fmt.Println("rev:", n)
		fmt.Println(buf.Bytes())
	}
}

// send generates random data of a given size and sends it to the server.
func send(size int) error {
	// Create a byte slice of the specified size and populate it with random data.
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	// Establish a TCP connection to the server.
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	// Send the size of the file to the server.
	err = binary.Write(conn, binary.LittleEndian, int64(size))
	if err != nil {
		log.Fatal(err)
	}

	// Send the actual file data to the server.
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		return err
	}

	// Print the number of bytes sent.
	fmt.Println("write byte :", n)
	return nil
}

func main() {
	// Start a goroutine that waits for 4 seconds before sending data to the server.
	go func() {
		time.Sleep(4 * time.Second)
		err := send(100)
		if err != nil {
			os.Exit(0)
		}
	}()

	// Start the file server.
	fs := &tcpServer{}
	fs.open()
}
