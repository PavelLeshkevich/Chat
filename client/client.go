package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	PORT       = ":8000"
	END_STRING = '\n'
)

func main() {
	conn, err := net.Dial("tcp", PORT)
	if err != nil {
		log.Fatal("Could not connect to the server", err)
	}
	defer conn.Close()
	connection := conn

	fmt.Print("Please entry your name: ")

	go func() {
		reader := bufio.NewReader(connection)
		for {
			msg, err := reader.ReadString(END_STRING)
			if err != nil {
				log.Fatal("Server Down! GoodBye!")
				break
			}
			fmt.Println(string(msg))
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadBytes(END_STRING)
		if err != nil {
			log.Println("Could not get message: ", err)
		}
		conn.Write(msg)
	}
}
