package handlers

import (
	"fmt"
	"strings"

	"github.com/kasirocswell/gomudserver/models"
)

func HandleCommand(user *models.User, command string, universe *models.Universe, args ...string) {
	if user == nil {
		fmt.Println("DEBUG: User is nil.")
		return
	}

	switch strings.ToLower(command) {
	case "equip":
		if len(args) == 0 {
			user.Writer.WriteString("Equip what?\n")
		} else {
			itemID := args[0]
			err := user.EquipItem(itemID)
			if err != nil {
				user.Writer.WriteString(err.Error() + "\n")
			} else {
				user.Writer.WriteString("Item equipped successfully.\n")
			}
		}
	case "unequip":
		if len(args) == 0 {
			user.Writer.WriteString("Unequip what?\n")
		} else {
			itemName := args[0]
			err := user.UnequipItem(itemName)
			if err != nil {
				user.Writer.WriteString(err.Error() + "\n")
			} else {
				user.Writer.WriteString("Item unequipped successfully.\n")
			}
		}

	case "status":
		user.Writer.WriteString("Character Status:\n")
		user.Writer.WriteString(fmt.Sprintf("Health: %d/%d\n", user.Health, user.MaxHealth))
		user.Writer.WriteString(fmt.Sprintf("Energy: %d/%d\n", user.Energy, user.MaxEnergy))
		user.Writer.WriteString(fmt.Sprintf("Current Weight: %d/%d\n", user.CurrentWeight, user.MaxCarryWeight))

		if len(user.StatusEffects) > 0 {
			user.Writer.WriteString("Status Effects:\n")
			for _, effect := range user.StatusEffects {
				user.Writer.WriteString(fmt.Sprintf(" - %s\n", effect))
			}
		} else {
			user.Writer.WriteString("No current status effects.\n")
		}

	case "talk":
		if len(args) == 0 {
			user.Writer.WriteString("Talk to whom?\n")
		} else {
			npcName := args[0]
			models.InteractWithNPC(user, npcName)
		}
	case "look":
		if user.Room != nil {
			user.Writer.WriteString(user.Room.Description(*universe))
		} else {
			user.Writer.WriteString("You are in an unknown location.\n")
		}

	case "n", "north", "s", "south", "e", "east", "w", "west", "enter", "exit":
		move(user, strings.ToLower(command), universe)
	case "inventory", "inv":
		// Handle inventory command
		inventoryList := user.ListInventory()
		if len(inventoryList) == 0 {
			user.Writer.WriteString("Your inventory is empty.\n")
		} else {
			user.Writer.WriteString("Your inventory contains:\n")
			for _, item := range inventoryList {
				user.Writer.WriteString(item + "\n")
			}
			user.Writer.WriteString(fmt.Sprintf("Total Weight: %d/%d\n", user.CurrentWeight, user.MaxCarryWeight))
		}
	default:
		user.Writer.WriteString("Unknown command.\n")
	}
	user.Writer.Flush()
}

func move(user *models.User, direction string, universe *models.Universe) {
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
	case "enter":
		nextRoom = user.Room.Enter
	case "exit":
		nextRoom = user.Room.Exit
	}

	if nextRoom == nil {
		user.Writer.WriteString("There's no exit in that direction.\n")
	} else {
		user.Room = nextRoom
		user.Writer.WriteString("You move " + direction + " to " + nextRoom.Name + ".\n")

		// Update the room description with the new location
		locationDescription := getLocationDescription(user, universe)

		user.Writer.WriteString(locationDescription + "\n")
	}
	user.Writer.Flush()
}

func getLocationDescription(user *models.User, universe *models.Universe) string {
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

	// Use galaxyName, planetName, and cityName in the description
	description := fmt.Sprintf("Galaxy: %s\nPlanet: %s\nCity: %s\n", galaxyName, planetName, cityName)
	description += user.Room.Description(*universe) // Append room description
	return description
}
