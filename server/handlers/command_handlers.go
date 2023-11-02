package handlers

import (
	"fmt"
	"strings"

	"github.com/kasirocswell/gomudserver/models"
)

// HandleCommand processes user commands and takes appropriate actions.
func HandleCommand(user *models.User, command string) {
	if user == nil {
		fmt.Println("DEBUG: User is nil.")
		return
	}

	switch strings.ToLower(command) {
	case "look":
		if user.Room == nil {
			user.Writer.WriteString("You are in an unknown location.\n")
		} else {
			// Construct the location description based on the user's current position
			locationDescription := getLocationDescription(user)
			user.Writer.WriteString(locationDescription + "\n")
		}
	case "n", "north", "s", "south", "e", "east", "w", "west", "enter", "exit":
		move(user, strings.ToLower(command))
	default:
		user.Writer.WriteString("Unknown command.\n")
	}
	user.Writer.Flush()
}

// move is a helper function to handle movement commands.
func move(user *models.User, direction string) {
	if user.Room == nil {
		user.Writer.WriteString("You are in an unknown location and cannot move.\n")
		user.Writer.Flush()
		return
	}

	var nextRoom *models.Room
	switch direction {
	case "n", "north":
		nextRoom = user.Room.N
	case "s", "south":
		nextRoom = user.Room.S
	case "e", "east":
		nextRoom = user.Room.E
	case "w", "west":
		nextRoom = user.Room.W
	}

	if nextRoom == nil {
		user.Writer.WriteString("There's no exit in that direction.\n")
	} else {
		user.Room = nextRoom
		user.Writer.WriteString("You move " + direction + " to " + user.Room.Name + ".\n")
		// Update the room description with the new location
		locationDescription := getLocationDescription(user)
		user.Writer.WriteString(locationDescription + "\n")
	}
	user.Writer.Flush()
}

// getLocationDescription constructs a string that describes the user's current location.
func getLocationDescription(user *models.User) string {
	galaxyName := "Unknown Galaxy"
	planetName := "Unknown Planet"
	cityName := "Unknown City"

	if user.Galaxy != nil {
		galaxyName = user.Galaxy.Name
	}
	if user.Planet != nil {
		planetName = user.Planet.Name
	}
	if user.City != nil {
		cityName = user.City.Name
	}

	// Call the Description method with the current location names
	return user.Room.Description(galaxyName, planetName, cityName)
}
