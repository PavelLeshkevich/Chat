package server

import "net"

type Chat struct {
	clients  []*Client
	connect  chan net.Conn
	outgoing chan string
}

func CreateChat() *Chat {
	chat := &Chat{
		clients:  make([]*Client, 0),
		connect:  make(chan net.Conn),
		outgoing: make(chan string),
	}

	go chat.Listen()

	return chat
}

func (chat *Chat) Connect(conn net.Conn) {
	chat.connect <- conn
}

func (chat *Chat) Listen() {
	for {
		select {
			case conn := <- chat.connect:
				chat.Join(conn)
			case msg := <- chat.outgoing:
				chat.Broadcast(msg)
		}
	}
}

func (chat *Chat) Join(conn net.Conn) {
	client := CreateClient(conn)
	chat.clients = append(chat.clients, client)
	go func() {
		for {
			chat.outgoing <- <- client.incoming
		}
	}()
}

func (chat *Chat) Broadcast(data string) {
	chat.UpdateClientsList()
	for _, client := range chat.clients {
		client.outgoing <- data
	}
}

func (chat * Chat) UpdateClientsList() {
	it := 0 // output index
	for _, client := range chat.clients {
		if client.status == true {
			chat.clients[it] = client
			it++
		}
	}
	chat.clients = chat.clients[:it]
}
