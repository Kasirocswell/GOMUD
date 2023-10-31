// universe.go
package models

import (
	"fmt"
)

// Universe structure
type Universe struct {
	Galaxies map[string]*Galaxy
}

func CreateUniverse() *Universe {
	// Define common rooms
	commonRooms := map[string]*Room{
		"Barracks":    CreateCommonRoom("Barracks", "A military facility in %s, %s where soldiers train and rest."),
		"Hospital":    CreateCommonRoom("Hospital", "A medical facility in %s, %s where the injured are treated."),
		"Seedy Bar":   CreateCommonRoom("Seedy Bar", "A dimly lit bar in %s, %s, filled with pilots and shady characters."),
		"Local Shop":  CreateCommonRoom("Local Shop", "A shop in %s, %s where entrepreneurs sell various goods."),
		"Town Square": CreateCommonRoom("Town Square", "The heart of %s, %s, bustling with activity."),
	}

	// Define unique rooms for each city
	uniqueRoomsForNeoTokyo := map[string]*Room{
		"NeoTokyo Tower": CreateUniqueRoom("NeoTokyo Tower", "A towering skyscraper offering a panoramic view of NeoTokyo."),
	}

	uniqueRoomsForCyberLisbon := map[string]*Room{
		"Cyber Cafe": CreateUniqueRoom("Cyber Cafe", "A popular cafe in CyberLisbon known for its digital espresso."),
	}

	uniqueRoomsForSiliconParis := map[string]*Room{
		"Silicon Museum": CreateUniqueRoom("Silicon Museum", "A museum in SiliconParis showcasing the history of technology."),
	}

	uniqueRoomsForDigitalDelhi := map[string]*Room{
		"Digital Market": CreateUniqueRoom("Digital Market", "A bustling market in DigitalDelhi with various digital goods and services."),
	}

	// Cities for Digitalis
	uniqueRoomsForDigitalisCity1 := map[string]*Room{
		"Digitalis Park": CreateUniqueRoom("Digitalis Park", "A serene park in the heart of the city."),
	}

	uniqueRoomsForDigitalisCity2 := map[string]*Room{
		"Digitalis Library": CreateUniqueRoom("Digitalis Library", "A vast library with digital archives."),
	}

	// Cities for NetNeptune
	uniqueRoomsForNeptuneCity1 := map[string]*Room{
		"Neptune Beach": CreateUniqueRoom("Neptune Beach", "A beautiful beach with glowing sands."),
	}

	uniqueRoomsForNeptuneCity2 := map[string]*Room{
		"Neptune Observatory": CreateUniqueRoom("Neptune Observatory", "A place to observe the vastness of the universe."),
	}

	// Create cities for each planet
	citiesForTechterra := map[string]*City{
		"NeoTokyo":    createCity("NeoTokyo", commonRooms, uniqueRoomsForNeoTokyo),
		"CyberLisbon": createCity("CyberLisbon", commonRooms, uniqueRoomsForCyberLisbon),
	}

	citiesForSiliconSphere := map[string]*City{
		"SiliconParis": createCity("SiliconParis", commonRooms, uniqueRoomsForSiliconParis),
		"DigitalDelhi": createCity("DigitalDelhi", commonRooms, uniqueRoomsForDigitalDelhi),
	}

	citiesForDigitalis := map[string]*City{
		"DigitalisCity1": createCity("DigitalisCity1", commonRooms, uniqueRoomsForDigitalisCity1),
		"DigitalisCity2": createCity("DigitalisCity2", commonRooms, uniqueRoomsForDigitalisCity2),
	}

	citiesForNetNeptune := map[string]*City{
		"NeptuneCity1": createCity("NeptuneCity1", commonRooms, uniqueRoomsForNeptuneCity1),
		"NeptuneCity2": createCity("NeptuneCity2", commonRooms, uniqueRoomsForNeptuneCity2),
	}

	// Create planets with their own set of cities
	planets := map[string]*Planet{
		"Techterra":     CreatePlanet("Techterra", citiesForTechterra),
		"SiliconSphere": CreatePlanet("SiliconSphere", citiesForSiliconSphere),
		"Digitalis":     CreatePlanet("Digitalis", citiesForDigitalis),
		"NetNeptune":    CreatePlanet("NetNeptune", citiesForNetNeptune),
	}

	// Create galaxies with planets
	galaxyNames := []string{"CyberCluster", "DigitalDomain", "NetNebula", "TechTwilight"}
	galaxies := make(map[string]*Galaxy)
	for _, galaxyName := range galaxyNames {
		galaxies[galaxyName] = createGalaxy(galaxyName, planets)
	}

	// Create universe with galaxies
	return &Universe{
		Galaxies: galaxies,
	}
}

func (u *Universe) GetGalaxy(name string) *Galaxy {
	galaxy, ok := u.Galaxies[name]
	if !ok {
		fmt.Printf("DEBUG: Galaxy %s not found.\n", name)
	}
	return galaxy
}

func (u *Universe) GetPlanet(galaxyName, planetName string) *Planet {
	galaxy := u.GetGalaxy(galaxyName)
	if galaxy != nil {
		planet, ok := galaxy.Planets[planetName]
		if !ok {
			fmt.Printf("DEBUG: Planet %s not found in Galaxy %s.\n", planetName, galaxyName)
		}
		return planet
	}
	return nil
}

func (u *Universe) GetCity(galaxyName, planetName, cityName string) *City {
	planet := u.GetPlanet(galaxyName, planetName)
	if planet != nil {
		city, ok := planet.Cities[cityName]
		if !ok {
			fmt.Printf("DEBUG: City %s not found on Planet %s in Galaxy %s.\n", cityName, planetName, galaxyName)
		}
		return city
	}
	return nil
}

// GetRoom retrieves a room by its name from a specified city in a specified planet in a specified galaxy.
func (u *Universe) GetRoom(galaxyName, planetName, cityName, roomName string) *Room {
	city := u.GetCity(galaxyName, planetName, cityName)
	if city != nil {
		if room, ok := city.Rooms[roomName]; ok {
			fmt.Printf("DEBUG: Galaxy %s not found.\n", roomName)
			return room
		}
	}
	fmt.Printf("DEBUG: Galaxy %s not found.\n", roomName)

	return nil
}
