package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kasirocswell/gomudserver/models"
)

// var AllNPCs = make(map[string]*models.NPC)

func main() {
	// Create the universe
	universe, err := models.CreateUniverse()
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

	// Set the path to your dialogue nodes data
	dialogueNodesFolderPath := "data/dialogue/"

	// Load dialogue nodes
	dialogueNodes, err := models.LoadDialogueNodes(dialogueNodesFolderPath)
	if err != nil {
		fmt.Println("Error loading dialogue nodes:", err)
		return
	}

	err = models.LoadQuests("data/quests")
	if err != nil {
		fmt.Println("Failed to load quests:", err)
		return
	}

	err = models.LoadWeapons("data/weapons")
	if err != nil {
		fmt.Printf("Failed to load weapons: %v", err)
	}

	err = models.LoadArmor("data/armor")
	if err != nil {
		log.Fatalf("Failed to load armor: %v", err)
	}

	err = models.LoadItems("data/items")
	if err != nil {
		log.Printf("Failed to load items: %v", err)
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
		go models.HandleNewConnection(conn, universe, dialogueNodes)
	}
}
