package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func setupRoutes() {
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/ws", WsEndpoint)
}

func main() {
	log.Println("Go WebSockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Homepage")
	fmt.Fprintf(w, "Welcome")
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage() //read messagenya dari user
		if err != nil {
			log.Printf("FuncReader, %s \n", err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("FuncReader WriteMessage %s \n", err)
			return
		}

	}
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to ws endpoint")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("%s \n", err)
	}

	log.Println("Client Successfully connected.")

	reader(ws)
}
