// city.go
package models

// City structure
type City struct {
	Name  string
	Rooms map[string]*Room
}

// createCity creates a new city with the given name and common rooms.
func createCity(name string, commonRooms map[string]*Room) *City {
	cityRooms := make(map[string]*Room)
	for roomName, room := range commonRooms {
		cityRooms[roomName] = room
	}
	return &City{
		Name:  name,
		Rooms: cityRooms,
	}
}
