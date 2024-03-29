package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// TODO: write server program to handle concurrent client connections.
	server, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		connection, err := server.Accept()
		if err != nil {
			continue
		}

		go handleConn(connection)
	}
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
