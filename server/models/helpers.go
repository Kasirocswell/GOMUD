package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func linkRooms(rooms map[string]*Room) {
	for _, room := range rooms {
		fmt.Printf("Linking room: %s - %s\n", room.ID, room.Name)
		fmt.Printf("  NorthID: %s, SouthID: %s, EastID: %s, WestID: %s, EnterID: %s, ExitID: %s\n",
			room.NorthID, room.SouthID, room.EastID, room.WestID, room.EnterID, room.ExitID)

		// Link to the North room
		if northID := room.NorthID; northID != "" {
			fmt.Printf("  Attempting to link North to %s\n", northID)
			if northRoom, exists := rooms[northID]; exists {
				room.N = northRoom
				fmt.Printf("  Successfully linked North to %s\n", northID)
			} else {
				fmt.Printf("  Failed to link North to %s\n", northID)
			}
		}

		// Link to the South room
		if southID := room.SouthID; southID != "" {
			fmt.Printf("  Attempting to link South to %s\n", southID)
			if southRoom, exists := rooms[southID]; exists {
				room.S = southRoom
				fmt.Printf("  Successfully linked South to %s\n", southID)
			} else {
				fmt.Printf("  Failed to link South to %s\n", southID)
			}
		}

		// Link to the East room
		if eastID := room.EastID; eastID != "" {
			fmt.Printf("  Attempting to link East to %s\n", eastID)
			if eastRoom, exists := rooms[eastID]; exists {
				room.E = eastRoom
				fmt.Printf("  Successfully linked East to %s\n", eastID)
			} else {
				fmt.Printf("  Failed to link East to %s\n", eastID)
			}
		}

		// Link to the West room
		if westID := room.WestID; westID != "" {
			fmt.Printf("  Attempting to link West to %s\n", westID)
			if westRoom, exists := rooms[westID]; exists {
				room.W = westRoom
				fmt.Printf("  Successfully linked West to %s\n", westID)
			} else {
				fmt.Printf("  Failed to link West to %s\n", westID)
			}
		}

		// Link to the Enter room
		if enterID := room.EnterID; enterID != "" {
			fmt.Printf("  Attempting to link Enter to %s\n", enterID)
			if enterRoom, exists := rooms[enterID]; exists {
				room.Enter = enterRoom
				fmt.Printf("  Successfully linked Enter to %s\n", enterID)
			} else {
				fmt.Printf("  Failed to link Enter to %s\n", enterID)
			}
		}

		// Link to the Exit room
		if exitID := room.ExitID; exitID != "" {
			fmt.Printf("  Attempting to link Exit to %s\n", exitID)
			if exitRoom, exists := rooms[exitID]; exists {
				room.Exit = exitRoom
				fmt.Printf("  Successfully linked Exit to %s\n", exitID)
			} else {
				fmt.Printf("  Failed to link Exit to %s\n", exitID)
			}
		}
	}
}

func (u *Universe) GetGalaxy(name string) *Galaxy {
	if galaxy, ok := u.Galaxies[name]; ok {
		return galaxy
	}
	return nil
}

func (u *Universe) GetPlanet(galaxyName, planetName string) *Planet {
	galaxy := u.GetGalaxy(galaxyName)
	if galaxy != nil {
		if planet, ok := galaxy.Planets[planetName]; ok {
			return planet
		}
	}
	return nil
}

// GetCity retrieves a city by its name from a specified planet in a specified galaxy.
func (u *Universe) GetCity(galaxyName, planetName, cityName string) *City {
	if galaxy, ok := u.Galaxies[galaxyName]; ok {
		if planet, ok := galaxy.Planets[planetName]; ok {
			if city, ok := planet.Cities[cityName]; ok {
				return city
			}
			fmt.Printf("City %s not found in Planet %s\n", cityName, planetName)
		} else {
			fmt.Printf("Planet %s not found in Galaxy %s\n", planetName, galaxyName)
		}
	} else {
		fmt.Printf("Galaxy %s not found\n", galaxyName)
	}
	return nil
}

// GetRoom retrieves a room by its name from a specified city in a specified planet in a specified galaxy.
func (u *Universe) GetRoom(galaxyName, planetName, cityName, roomName string) *Room {
	if galaxy, ok := u.Galaxies[galaxyName]; ok {
		if planet, ok := galaxy.Planets[planetName]; ok {
			if city, ok := planet.Cities[cityName]; ok {
				if room, ok := city.Rooms[roomName]; ok {
					return room
				}
				fmt.Printf("Room %s not found in City %s\n", roomName, cityName)
			} else {
				fmt.Printf("City %s not found in Planet %s\n", cityName, planetName)
			}
		} else {
			fmt.Printf("Planet %s not found in Galaxy %s\n", planetName, galaxyName)
		}
	} else {
		fmt.Printf("Galaxy %s not found\n", galaxyName)
	}
	return nil
}

func Contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

var AllNPCs = make(map[string]*NPC)

func LoadNPCs(folderPath string, universe *Universe) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(folderPath, file.Name())
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		var npcs []NPC
		if err := json.Unmarshal(fileData, &npcs); err != nil {
			return err
		}

		for _, npc := range npcs {
			npcCopy := npc             // Create a copy of the NPC to avoid pointer issues
			AllNPCs[npc.ID] = &npcCopy // Add the NPC copy to the global map

			if room := findRoomByNPCLocation(npc.LocationID, universe); room != nil {
				// Check if NPC is already in the room
				if !Contains(room.NPCs, npc.ID) {
					room.NPCs = append(room.NPCs, npc.ID) // Assign the NPC ID to the room
				}
			}
		}
	}

	return nil
}

func findRoomByNPCLocation(roomID string, universe *Universe) *Room {
	for _, galaxy := range universe.Galaxies {
		for _, planet := range galaxy.Planets {
			for _, city := range planet.Cities {
				if room, ok := city.Rooms[roomID]; ok {
					return room
				}
			}
		}
	}
	return nil
}

type DialogueNodesContainer struct {
	Nodes []DialogueNode `json:"nodes"`
}

func LoadDialogueNodes(folderPath string) (map[string]DialogueNode, error) {
	dialogueNodes := make(map[string]DialogueNode)

	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		filePath := filepath.Join(folderPath, file.Name())
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		var container DialogueNodesContainer
		if err := json.Unmarshal(fileData, &container); err != nil {
			return nil, err
		}

		for _, node := range container.Nodes {
			dialogueNodes[node.ID] = node
		}
	}

	return dialogueNodes, nil
}
