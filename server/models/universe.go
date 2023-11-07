package models

// Import necessary packages

// Universe structure
type Universe struct {
	Galaxies map[string]*Galaxy
}

// NewUniverse creates a new Universe instance.
func NewUniverse() *Universe {
	return &Universe{
		Galaxies: make(map[string]*Galaxy),
	}
}

// CreateUniverse initializes the universe with its galaxies, planets, and cities.
func CreateUniverse() *Universe {
	// Define common rooms
	commonRooms := InitializeCommonRooms()

	// Initialize unique rooms for each city
	uniqueRoomsForNeoTokyo := InitializeUniqueRooms("NeoTokyo")
	uniqueRoomsForCyberLisbon := InitializeUniqueRooms("CyberLisbon")
	uniqueRoomsForSiliconParis := InitializeUniqueRooms("SiliconParis")
	uniqueRoomsForDigitalDelhi := InitializeUniqueRooms("DigitalDelhi")
	uniqueRoomsForDigitalisCity1 := InitializeUniqueRooms("DigitalisCity1")
	uniqueRoomsForDigitalisCity2 := InitializeUniqueRooms("DigitalisCity2")
	uniqueRoomsForNeptuneCity1 := InitializeUniqueRooms("NeptuneCity1")
	uniqueRoomsForNeptuneCity2 := InitializeUniqueRooms("NeptuneCity2")

	// Create cities with common and unique rooms for each planet
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
	galaxies := map[string]*Galaxy{
		"CyberCluster":  createGalaxy("CyberCluster", map[string]*Planet{"Techterra": planets["Techterra"]}),
		"SiliconDomain": createGalaxy("SiliconDomain", map[string]*Planet{"SiliconSphere": planets["SiliconSphere"]}),
		"DataDell":      createGalaxy("DataDell", map[string]*Planet{"Digitalis": planets["Digitalis"]}),
		"NeptuneNest":   createGalaxy("NeptuneNest", map[string]*Planet{"NetNeptune": planets["NetNeptune"]}),
	}

	// Link unique rooms within cities if necessary
	// For example, if you have a secret passage from the Secret Lab to the Armory in NeoTokyo
	// citiesForTechterra["NeoTokyo"].Rooms["Secret Lab"].Enter = citiesForTechterra["NeoTokyo"].Rooms["Armory"]
	// citiesForTechterra["NeoTokyo"].Rooms["Armory"].Exit = citiesForTechterra["NeoTokyo"].Rooms["Secret Lab"]

	// ... Link unique rooms for other cities similarly

	// Create universe with galaxies
	return &Universe{
		Galaxies: galaxies,
	}
}

// GetGalaxy retrieves a galaxy by its name from the universe.
func (u *Universe) GetGalaxy(name string) *Galaxy {
	if galaxy, ok := u.Galaxies[name]; ok {
		return galaxy
	}
	return nil
}

// GetPlanet retrieves a planet by its name from a specified galaxy.
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
	planet := u.GetPlanet(galaxyName, planetName)
	if planet != nil {
		if city, ok := planet.Cities[cityName]; ok {
			return city
		}
	}
	return nil
}

// GetRoom retrieves a room by its name from a specified city in a specified planet in a specified galaxy.
func (u *Universe) GetRoom(galaxyName, planetName, cityName, roomName string) *Room {
	city := u.GetCity(galaxyName, planetName, cityName)
	if city != nil {
		if room, ok := city.Rooms[roomName]; ok {
			return room
		}
	}
	return nil
}
