package main

import (
	"io"
	"log"
	"net"
)

const (
	PORT = ":8080"
	TYPE = "tcp"
)

func main() {
	listen, err := net.Listen(TYPE, PORT)
	if err != nil {
		log.Fatalf("Error strating server: %s", err)
	}
	// close listener
	defer listen.Close()
	log.Printf("Server started on port %s", PORT)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Error accepting: %s", err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	// incoming request
	_, err := io.Copy(conn, conn)
	if err != nil {
		log.Fatal(err)
	}

}
