package main

import (
	"fmt"
	"net"

	"github.com/kasirocswell/gomudserver/handlers"
	"github.com/kasirocswell/gomudserver/models"
)

var AllNPCs = make(map[string]*models.NPC)

func main() {
	// Create the universe
	universe, err := models.CreateUniverse(AllNPCs)
	if err != nil {
		fmt.Println("Error creating universe:", err)
		return
	}

	// Load NPCs into the universe
	err = models.LoadNPCs("data/npc/", universe)
	if err != nil {
		fmt.Println("Error loading NPCs:", err)
		return
	}

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

		// Pass the universe to the connection handler
		go handlers.HandleNewConnection(conn, universe)
	}
}
