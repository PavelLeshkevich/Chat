package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	PORT = 8000
)

var (
	connection net.Conn
	reader     *bufio.Reader
	writer     *bufio.Writer
)

func main() {
	conn, err := net.Dial("tcp", ":"+string(PORT))
	if err != nil {
		log.Fatal("Could not connect to the server", err)
		return
	}
	defer conn.Close()
	connection = conn

	reader = bufio.NewReader(connection)
	writer = bufio.NewWriter(connection)

	writer.WriteString("Please write you name :)")
	writer.Flush()

	go func() {
		for {
			data, err := reader.ReadString('\n')
			if err != nill {
				log.Fatal("Server down ):")
				break
			}
			fmt.Println(data)
		}
	}()
}
