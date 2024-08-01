package main

import (
	server "Priente/internal"
	"log"
	"net/http"
)

func main() {
	// Create a file server to serve static files from the "static" directory as the root directory
	fileServer := http.FileServer(http.Dir("./src"))

	// Handle requests to the root by serving files from the "static" directory
	http.Handle("/", fileServer)

	// Initialize the WebSocket server
	wsServer := server.NewWebSocketServer()

	// Handle WebSocket requests
	http.HandleFunc("/ws", wsServer.HandleConnections)
	go wsServer.HandleMessages()

	// Start the server on port 8080
	log.Println("Server is listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed: ", err)
	}
}
