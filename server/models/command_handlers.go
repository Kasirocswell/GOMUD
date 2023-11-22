package models

import (
	"fmt"
	"strings"
)

func HandleCommand(user *User, command string, universe *Universe, dialogueNodes map[string]DialogueNode, args ...string) {
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
		if len(user.Room.NPCs) == 0 {
			user.Writer.WriteString("There's no one to talk to here.\n")
		} else {
			// Start dialogue with the first NPC in the room
			npcID := user.Room.NPCs[0]
			npc, exists := AllNPCs[npcID]
			if exists && npc.DialogueStartNode != "" {
				InteractWithNPC(user, npc, dialogueNodes)
				// The dialogue interaction is handled within InteractWithNPC.
				// Once this function returns, the player is back in the game state.
			} else {
				user.Writer.WriteString("There's no one to talk to here.\n")
			}
		}
	case "drop":
		if len(args) == 0 {
			user.Writer.WriteString("Drop what?\n")
		} else {
			itemName := strings.Join(args, " ") // Assuming item names can have spaces
			item, err := user.RemoveItemFromInventory(itemName)
			if err != nil {
				user.Writer.WriteString(err.Error() + "\n")
			} else {
				user.Room.Items = append(user.Room.Items, item)
				user.Writer.WriteString(fmt.Sprintf("You dropped %s.\n", itemName))
			}
		}
		user.Writer.Flush()

		user.Writer.Flush()
	case "quests":
		user.ListQuests()
	case "look":
		if user.Room != nil {
			// Correct argument (the global NPC map)
			user.Writer.WriteString(user.Room.Description(universe, AllNPCs))

		} else {
			user.Writer.WriteString("You are in an unknown location.\n")
		}

	case "n", "north", "s", "south", "e", "east", "w", "west", "enter", "exit":
		move(user, strings.ToLower(command), universe)
	case "inventory", "inv":
		// Handle inventory command
		inventoryList := user.ListInventory()
		if len(inventoryList) == 0 && user.Equipment.IsEmpty() {
			user.Writer.WriteString("Your inventory and equipment are empty.\n")
		} else {
			if len(inventoryList) > 0 {
				user.Writer.WriteString("Your inventory contains:\n")
				for _, item := range inventoryList {
					user.Writer.WriteString(item + "\n")
				}
				user.Writer.WriteString(fmt.Sprintf("Total Weight: %d/%d\n", user.CurrentWeight, user.MaxCarryWeight))
			} else {
				user.Writer.WriteString("Your inventory is empty.\n")
			}

			user.Writer.WriteString("You are equipped with:\n")
			user.Equipment.ListEquipment(user.Writer)
		}
		user.Writer.Flush()

	default:
		user.Writer.WriteString("Unknown command.\n")
	}
	user.Writer.Flush()
}

func move(user *User, direction string, universe *Universe) {
	if user.Room == nil {
		user.Writer.WriteString("You are in an unknown location and cannot move.\n")
		user.Writer.Flush()
		return
	}

	var nextRoom *Room // Declare as a pointer
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
		user.Room = nextRoom // This is now a pointer-to-pointer assignment
		user.Writer.WriteString("You move " + direction + " to " + nextRoom.Name + ".\n")

		// Update the room description with the new location
		locationDescription := getLocationDescription(user, universe)

		user.Writer.WriteString(locationDescription + "\n")
	}
	user.Writer.Flush()
}

// getLocationDescription constructs a string that describes the user's current location.
func getLocationDescription(user *User, universe *Universe) string {
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

	// Construct the base description using galaxy, planet, and city names
	description := fmt.Sprintf("Galaxy: %s\nPlanet: %s\nCity: %s\n", galaxyName, planetName, cityName)

	// Append the room description, if the user is currently in a room
	if user.Room != nil {
		description += user.Room.Description(universe, AllNPCs)
	} else {
		description += "You are currently not in any specific room."
	}

	return description
}

func FindNPCByName(name string, room *Room) *NPC {
	for _, npcID := range room.NPCs {
		npc, exists := AllNPCs[npcID]
		if exists && strings.EqualFold(npc.Name, name) {
			return npc
		}
	}
	return nil
}
