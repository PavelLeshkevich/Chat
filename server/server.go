package server

import (
	"bufio"
	"fmt"
	"net"
)

const (
	END_STRING = '\n'
)

type Client struct {
	name       string
	conn       net.Conn
	writer     *bufio.Writer
	reader     *bufio.Reader
	incoming   chan string
	outgoing   chan string
	disconnect chan bool
	status     bool // true connected, false otherwise
}

func CreateClient(conn net.Conn) *Client {
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	client := &Client{
		name:     "user",
		conn:     conn,
		writer:   writer,
		outgoing: make(chan string),
		reader:   reader,
		incoming: make(chan string),
		disconnect: make(chan bool),
		status:   true,
	}

	go client.Write()
	go client.Read()

	return client
}

func (client *Client) Write() {
	for {
		select {
		case <-client.disconnect:
			client.status = false
			break
		case msg := <- client.outgoing:
			client.writer.WriteString(msg)
			client.writer.Flush()
		}
	}
}

func (client *Client) Read() {
	for {
		msg, err := client.reader.ReadString(END_STRING)
		if err != nil {
			client.incoming <- fmt.Sprintf("%s disconnected", client.name)
			client.disconnect <- true
			client.conn.Close()
			break
		}
		client.incoming <- fmt.Sprintf("%s: %s", client.name, msg)
	}
}
