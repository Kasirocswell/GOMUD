package models

import (
	"fmt"
)

const (
	GalaxiesJSONPath = "data/galaxies.json" // Path to your galaxies JSON
	PlanetsJSONPath  = "data/planets.json"  // Path to your planets JSON
	CitiesJSONPath   = "data/cities.json"   // Path to your folder containing city JSON files
	RoomsFolderPath  = "data/rooms/"
	npcFilePath      = "data/npc"
)

// Universe structure
type Universe struct {
	Galaxies map[string]*Galaxy
	NPCs     map[string]*NPC
}

// NewUniverse creates a new Universe instance.
func NewUniverse() *Universe {
	return &Universe{
		Galaxies: make(map[string]*Galaxy),
	}
}

func CreateUniverse(npcMap map[string]*NPC) (*Universe, error) {
	universe := NewUniverse()

	// Load Galaxies
	galaxies, err := LoadGalaxies(GalaxiesJSONPath)
	if err != nil {
		return nil, fmt.Errorf("error loading galaxies: %w", err)
	}

	// Load Planets
	planets, err := LoadPlanets(PlanetsJSONPath)
	if err != nil {
		return nil, fmt.Errorf("error loading planets: %w", err)
	}

	// Load Cities
	cities, err := LoadCities(CitiesJSONPath)
	if err != nil {
		return nil, fmt.Errorf("error loading cities: %w", err)
	}

	// Load and Assign Rooms
	roomsByCity, err := LoadRoomsFromFolder(RoomsFolderPath)
	if err != nil {
		return nil, fmt.Errorf("error loading rooms: %w", err)
	}

	// Process the loaded data
	for _, galaxy := range galaxies {
		galaxyMap := &Galaxy{
			ID:      galaxy.ID,
			Name:    galaxy.Name,
			Planets: make(map[string]*Planet),
		}
		universe.Galaxies[galaxy.ID] = galaxyMap

		for _, planet := range planets {
			if planet.GalaxyID == galaxy.ID {
				planetMap := &Planet{
					ID:       planet.ID,
					GalaxyID: planet.GalaxyID,
					Name:     planet.Name,
					Cities:   make(map[string]*City),
				}
				galaxyMap.Planets[planet.ID] = planetMap

				for _, city := range cities {
					if city.PlanetID == planet.ID {
						cityMap := &City{
							ID:       city.ID,
							PlanetID: city.PlanetID,
							Name:     city.Name,
							Rooms:    make(map[string]*Room),
						}
						planetMap.Cities[city.ID] = cityMap

						if rooms, exists := roomsByCity[city.ID]; exists {
							for _, roomData := range rooms {
								roomMap := &Room{
									ID:              roomData.ID,
									CityID:          roomData.CityID,
									Name:            roomData.Name,
									DescriptionText: roomData.DescriptionText,
									Players:         make(map[*User]bool),
									NorthID:         roomData.NorthID,
									SouthID:         roomData.SouthID,
									EastID:          roomData.EastID,
									WestID:          roomData.WestID,
									EnterID:         roomData.EnterID,
									ExitID:          roomData.ExitID,
								}
								cityMap.Rooms[roomData.ID] = roomMap
							}
							// Link rooms within the city
							linkRooms(cityMap.Rooms)
						}
					}
				}
			}
		}
	}

	// Load NPCs
	err = LoadNPCs(npcFilePath, universe)
	if err != nil {
		return nil, fmt.Errorf("error loading NPCs: %w", err)
	}

	return universe, nil
}
