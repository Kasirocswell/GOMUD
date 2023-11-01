package main

import (
	"fmt"
	"net"

	"github.com/kasirocswell/gomudserver/handlers"
	"github.com/kasirocswell/gomudserver/models"
)

func main() {
	// Create the universe
	universe := models.CreateUniverse()

	// Start server
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("MUD server started on localhost:4000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handlers.HandleNewConnection(conn, universe)
	}
}
