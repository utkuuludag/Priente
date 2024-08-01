package main

import (
	server "Priente/internal"
	"log"
	"net/http"
)

func main() {
	wsServer := server.NewWebSocketServer()

	http.HandleFunc("/ws", wsServer.HandleConnections)

	go wsServer.HandleMessages()

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
