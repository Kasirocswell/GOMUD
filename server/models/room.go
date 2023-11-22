package models

import (
	"fmt"
	"strings"
)

// Room structure
type Room struct {
	ID              string
	CityID          string
	Name            string
	DescriptionText string
	Players         map[*User]bool
	Items           []Item

	// IDs for adjacent rooms
	NorthID string
	SouthID string
	EastID  string
	WestID  string
	EnterID string
	ExitID  string

	// Pointers to adjacent rooms
	N     *Room
	S     *Room
	E     *Room
	W     *Room
	Enter *Room
	Exit  *Room

	NPCs []string
}

func (r *Room) Description(universe *Universe, npcMap map[string]*NPC) string {
	description := fmt.Sprintf("Location: %s\n\n%s", r.Name, r.DescriptionText)

	// Add NPCs present in the room to the description
	if len(r.NPCs) > 0 {
		npcNames := make([]string, 0, len(r.NPCs))
		for _, npcID := range r.NPCs {
			if npc, exists := npcMap[npcID]; exists {
				npcNames = append(npcNames, npc.Name)
			}
		}
		description += "\nNPCs here: " + strings.Join(npcNames, ", ")
	} else {
		description += "\nThere are no NPCs here."
	}

	// Add information about available exits
	exits := []string{}
	if r.N != nil {
		exits = append(exits, "north")
	}
	if r.S != nil {
		exits = append(exits, "south")
	}
	if r.E != nil {
		exits = append(exits, "east")
	}
	if r.W != nil {
		exits = append(exits, "west")
	}
	if r.Enter != nil {
		exits = append(exits, "enter")
	}
	if r.Exit != nil {
		exits = append(exits, "exit")
	}

	if len(exits) > 0 {
		description += "\nExits: " + strings.Join(exits, ", ")
	} else {
		description += "\nThere are no visible exits."
	}

	return description
}

// Additional functions for creating unique rooms or other utilities can be added here.
