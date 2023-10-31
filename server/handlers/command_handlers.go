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
		} else if user.City == nil || user.Planet == nil {
			user.Writer.WriteString("Unable to determine your exact location.\n")
		} else {
			user.Writer.WriteString(user.Room.Description(user.City.Name, user.Planet.Name) + "\n")
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
		return
	}

	if exit, ok := user.Room.Exits[direction]; ok {
		user.Room = exit
		if user.Room == nil {
			user.Writer.WriteString("You move to an unknown location.\n")
			return
		}
		user.Writer.WriteString("You move " + direction + " to " + user.Room.Name + ".\n")
		if user.City != nil && user.Planet != nil {
			user.Writer.WriteString(user.Room.Description(user.City.Name, user.Planet.Name) + "\n")
		} else {
			user.Writer.WriteString("Unable to determine your exact location after moving.\n")
		}
	} else {
		user.Writer.WriteString("There's no exit in that direction.\n")
	}
	user.Writer.Flush()
}
