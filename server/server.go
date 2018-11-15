package server

import (
	"bufio"
	"net"
)

type Client struct {
	name       string
	conn       net.Conn
	writer     *bufio.Writer
	reader     *bufio.Reader
	incoming   chan string
	outgoing   chan string
	disconnect chan bool
	status     int // 1 connected, 0 otherwise
}

func CreateClient(conn net.Conn) *Client {
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	client := &Client{
		name:       "user",
		conn:       conn,
		writer:     writer,
		outgoing:   make(chan string),
		reader:     reader,
		incoming:   make(chan string),
		disconnect: make(chan bool),
		status:     1,
	}

	go client.Write()
	go client.Read()

	return client
}
