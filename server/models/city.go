package models

// City structure
type City struct {
	Name  string
	Rooms map[string]*Room
}

// createCity creates a new city with the given name and common rooms.
// Now it also accepts a map of uniqueRooms.
func createCity(name string, commonRooms, uniqueRooms map[string]*Room) *City {
	cityRooms := make(map[string]*Room)
	// Add common rooms to the city
	for roomName, room := range commonRooms {
		cityRooms[roomName] = room
	}
	// Add unique rooms to the city
	for roomName, room := range uniqueRooms {
		cityRooms[roomName] = room
	}
	return &City{
		Name:  name,
		Rooms: cityRooms,
	}
}

// InitializeUniqueRooms creates a map of unique rooms for a specific city.
// This function should be expanded to include all unique rooms for each city.
func InitializeUniqueRooms(cityName string) map[string]*Room {
	uniqueRooms := make(map[string]*Room)
	switch cityName {
	case "NeoTokyo":
		// Define unique rooms for NeoTokyo
		secretLab := NewRoom("Secret Lab", "A hidden laboratory with advanced technology.", nil, nil, nil, nil, nil, nil)
		uniqueRooms["Secret Lab"] = secretLab
	case "CyberLisbon":
		// Define unique rooms for CyberLisbon
		ancientRuins := NewRoom("Ancient Ruins", "Ancient ruins from a bygone era.", nil, nil, nil, nil, nil, nil)
		uniqueRooms["Ancient Ruins"] = ancientRuins
		// Add cases for other cities
	}
	return uniqueRooms
}

// AddUniqueRoom adds a unique room to the city.
func (c *City) AddUniqueRoom(room *Room) {
	if c.Rooms == nil {
		c.Rooms = make(map[string]*Room)
	}
	c.Rooms[room.Name] = room
}

// LinkRoomsWithinCity connects rooms within the city.
func (c *City) LinkRoomsWithinCity(roomAName, roomBName, direction string) {
	roomA, roomAExists := c.Rooms[roomAName]
	roomB, roomBExists := c.Rooms[roomBName]
	if roomAExists && roomBExists {
		LinkRooms(roomA, roomB, direction)
	}
}

// LinkRooms is a utility function that connects two rooms bidirectionally.
func LinkRooms(roomA, roomB *Room, direction string) {
	switch direction {
	case "north":
		roomA.N = roomB
		roomB.S = roomA
	case "south":
		roomA.S = roomB
		roomB.N = roomA
	case "east":
		roomA.E = roomB
		roomB.W = roomA
	case "west":
		roomA.W = roomB
		roomB.E = roomA
	case "enter":
		roomA.Enter = roomB
		roomB.Exit = roomA
	case "exit":
		roomA.Exit = roomB
		roomB.Enter = roomA
	}
}

// Additional methods for the City struct can be added here.
