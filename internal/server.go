// internal/server.go
package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Message defines the structure of a message
type Message struct {
	Data string `json:"data"`
}

// WebSocketServer handles WebSocket connections
type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	mutex     sync.Mutex
	broadcast chan Message
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// NewWebSocketServer initializes a new WebSocket server
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan Message),
	}
}

// HandleConnections handles WebSocket requests from clients
func (server *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}
	defer conn.Close()

	server.addClient(conn)

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Read error: %v", err)
			server.removeClient(conn)
			break
		}
		server.broadcast <- msg
	}
}

// HandleMessages sends messages to all connected clients
func (server *WebSocketServer) HandleMessages() {
	for {
		msg := <-server.broadcast
		server.mutex.Lock()
		for client := range server.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Write error: %v", err)
				client.Close()
				delete(server.clients, client)
			}
		}
		server.mutex.Unlock()
	}
}

// addClient adds a new client to the list
func (server *WebSocketServer) addClient(conn *websocket.Conn) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.clients[conn] = true
}

// removeClient removes a client from the list
func (server *WebSocketServer) removeClient(conn *websocket.Conn) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	delete(server.clients, conn)
}
