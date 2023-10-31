// city.go
package models

// City structure
type City struct {
	Name  string
	Rooms map[string]*Room
}

func createCity(name string, commonRooms, uniqueRooms map[string]*Room) *City {
	cityRooms := make(map[string]*Room)
	for roomName, room := range commonRooms {
		cityRooms[roomName] = &Room{
			Name:                room.Name,
			DescriptionTemplate: room.DescriptionTemplate,
			Users:               make(map[*User]bool),
		}
	}
	for roomName, room := range uniqueRooms {
		cityRooms[roomName] = room
	}
	return &City{
		Name:  name,
		Rooms: cityRooms,
	}
}
