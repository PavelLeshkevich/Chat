package main

import (
	"log"
	"net"
)

const (
	PORT = ":8000"
)

func main() {
	log.Println("Server started")

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("Could not start server. Reason: ", err)
	}
}
