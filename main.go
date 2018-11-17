package main

import (
	"./server"
	"log"
	"net"
)

const (
	PORT = ":8000"
	END_STRING = "\n"
)

func main() {
	log.Println("Start server")

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("Error Listeing: ", err)
	}
	defer listener.Close()

	chat := server.CreateChat()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting: ",  err)
		}
		chat.Connect(conn)
	}
}
