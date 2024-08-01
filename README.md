# Priente

# Simple WebSocket Application in Go

This project is a simple WebSocket server implemented in GoLang. The server uses basic data structures and Go's concurrency features to provide several core functionalities, including managing WebSocket connections, broadcasting messages, and handling client connections.

## Features

- **WebSocket Management:** Uses the `github.com/gorilla/websocket` package to handle WebSocket connections.
- **Concurrency:** Utilizes Go's `sync.Mutex` and channels to safely manage concurrent operations.
- **Client Management:** Keeps track of connected clients using a map, with thread-safe access.
- **Message Broadcasting:** Relays messages received from one client to all other connected clients.
- **Graceful Disconnection:** Handles client disconnections gracefully.

## Requirements

- Go 1.16 or higher
- [Gorilla WebSocket package](https://github.com/gorilla/websocket)

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/utkuuludag/Priente.git
   cd Priente

2. **Initialize Go Modules:**

    go mod tidy

3. **Install Dependencies:**

    go get -u github.com/gorilla/websocket


- Project Structure
    - src/main.go: The entry point for the WebSocket server.
    - src/index.html: A simple client HTML page to test the WebSocket connection.
    - src/app.js: A simple JavaScript page to test WebSocket connection.
    - internal/server.go: Contains the WebSocket server logic.

- Running the Server
    1. Start the WebSocket Server:

        - Navigate to the project root directory and run: "go run src/main.go"
        - Alternatively, you can start the project using the main.exe located in the root directory: ./main.exe
        This will start the server on http://localhost:8080.
    2. Access the Web Interface:
        Open a web browser and go to http://localhost:8080 or http://localhost:8080/index.html to load the client interface.

- Testing the WebSocket
    To test the WebSocket server, open multiple browser windows or tabs and connect to the server using the provided client interface. You can send messages from any connected client, and all other clients should receive the broadcasted messages.