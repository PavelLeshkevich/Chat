package main

import (
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
	"os"
)

type (
	Msg struct {
		clientKey string
		text string
	}

	NewClientEvent struct {
		clientKey string
		msgChan chan Msg
 	}
)

const MAXBACKLOG = 100

var (
	dirPath string
	clientRequests = make(chan *NewClientEvent, 100)
	clientDisconnects = make(chan string, 100)
)

func IndexPage(w http.ResponseWriter, req *http.Request, filename string) {

	fp, err := os.Open(dirPath + "/" + filename)

	if err != nil {
		log.Println("Could not open file", err.Error())
		w.Write([]byte("500 internal server error"))
		return
	}

	defer fp.Close()

	_, err = io.Copy(w, fp)
	if err != nil {
		log.Println("Could not send file content", err.Error())
		w.Write([]byte("500 internal server error"))
		return
	}

}

func ChatServer(ws *websocket.Conn) {
	var lenBuf[5]byte

	ws.SetDeadline(timeout)

	msgChan := make(chan Msg, 100)
	clientKey := ws.RemoteAddr().String()
	clientRequests <- &NewClientEvent{clientKey, msgChan}
	defer func() { clientDisconnects <- clientKey}()

	for {
		err := ws.Read(msg)
	}
}

func router() {
	clients := make(map[string]chan Msg)

	for {
		select {
			case req := <-clientRequests:
				clients[req.clientKey] = req.msgChan
				log.Println("Websocket connected: " + req.clientKey)
			case clientKey := <-clientDisconnects:
				delete(clients, clientKey)
				log.Println("Websocket disconnected: " + clientKey)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: chatExample <dir>")

	}
	dirPath = os.Args[1]

	log.Println("Starting...")

	go router()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		IndexPage(w, req, "index.html")
	})
	http.HandleFunc("/index.js", func(w http.ResponseWriter, req *http.Request){
		IndexPage(w, req, "index.js")
	})
	http.Handle("/ws", websocket.Handler(ChatServer))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
