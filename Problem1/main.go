package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/big"
	"net"
)

const (
	PORT = ":8080"
	TYPE = "tcp"
)

type Client struct {
	Method string   `json:"method"`
	Number *float64 `json:"number"`
}

var (
	isPrime  = []byte(`{"method":"isPrime","prime":true}` + "\n")
	notPrime = []byte(`{"method":"isPrime","prime":false}` + "\n")
	notJson  = []byte(`malformed` + "\n")
)

func checkPrimeNumber(num float64) bool {
	result := big.NewInt(int64(num)).ProbablyPrime(0)
	return result
}

func checkJson(j []byte) []byte {
	var client Client
	err := json.Unmarshal(j, &client)
	if err != nil || client.Method != "isPrime" || client.Number == nil {
		return notJson
	} else if checkPrimeNumber(*client.Number) {
		return isPrime
	}
	return notPrime
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewScanner(conn)
	for reader.Scan() {
		result := checkJson(reader.Bytes())
		conn.Write(result)
	}
}

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
