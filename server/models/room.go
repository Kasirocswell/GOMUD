package models

import (
	"fmt"
	"strings"
)

// Room structure
type Room struct {
	Name            string
	DescriptionText string // Renamed from 'Description' to 'DescriptionText'
	Players         map[*User]bool
	N               *Room
	S               *Room
	E               *Room
	W               *Room
	Enter           *Room
	Exit            *Room
}

// NewRoom creates a new Room with the given attributes.
func NewRoom(name, descriptionText string, n, s, e, w, enter, exit *Room) *Room {
	return &Room{
		Name:            name,
		DescriptionText: descriptionText,
		Players:         make(map[*User]bool),
		N:               n,
		S:               s,
		E:               e,
		W:               w,
		Enter:           enter,
		Exit:            exit,
	}
}

// InitializeCommonRooms creates a map of common rooms.
func InitializeCommonRooms() map[string]*Room {
	barracks := NewRoom("Barracks", "Before you stands the imposing structure of the city's barracks, its fortified walls stretching upwards to meet the heavy battlements above. The constant hustle of guards moving in and out, some in the midst of training while others stand watch, conveys a sense of disciplined order. The clang of metal and shouted commands echo off the walls, a reminder of the city's readiness to defend its people. A heavy wooden door, reinforced with iron bands, serves as the main entrance, standing ajar during the day and securely bolted at night.", nil, nil, nil, nil, nil, nil)
	hospital := NewRoom("Hospital", "The city's hospital exudes a quiet air of solemnity, its white marble walls and clean, polished surfaces contrasting with the often grim realities faced within. The gentle hum of machinery and the soft footsteps of the medical staff provide a backdrop to the muted conversations of visitors and patients alike. The scent of antiseptics permeates the air, a testament to the ceaseless battle against infection and disease. Wide doors open into a series of treatment rooms, each equipped to deal with the myriad ailments that befall the city's denizens.", nil, nil, nil, nil, nil, nil)
	seedyBar := NewRoom("Seedy Bar", "The flickering neon sign of the Seedy Bar cuts through the murky twilight, casting a sordid glow on the establishment's dilapidated facade. The muffled sounds of raucous laughter and clinking glasses seep out from the cracks in the boarded-up windows, hinting at the dubious escapades unfolding within. A pungent mix of stale alcohol and sweat hangs heavy in the air, and the doorway stands as an unspoken challenge, daring the more adventurous—or desperate—souls to step inside.", nil, nil, nil, nil, nil, nil)
	localShop := NewRoom("Local Shop", "Nestled between the more imposing structures of the city, the Local Shop presents a quaint and welcoming facade. Its windows display an eclectic array of goods ranging from practical tools to exotic trinkets, each with its own tale of distant lands and bygone times. The warm light spilling out from the interior promises a respite from the bustling city streets, and the tinkling bell above the door sings a cheerful note with each new arrival, inviting passersby to explore the treasures within.", nil, nil, nil, nil, nil, nil)
	townSquare := NewRoom("Town Square", "The Town Square is the vibrant heart of the city, a spacious plaza filled with the constant hum of activity. Cobblestone paths crisscross the area, guiding the footsteps of merchants, travelers, and locals alike. The air carries the mingled scents of street food, the sounds of lively music, and the chatter of commerce. At the center, a historic fountain stands as a monument to the city's heritage, its waters a sparkling meeting point for social exchanges and quiet contemplation amidst the urban energy.", nil, nil, nil, nil, nil, nil)

	// Link rooms together
	barracks.E = hospital
	hospital.W = barracks
	hospital.N = seedyBar
	seedyBar.S = hospital
	seedyBar.E = localShop
	localShop.W = seedyBar
	localShop.N = townSquare
	townSquare.S = localShop

	// Return a map of the rooms
	return map[string]*Room{
		"Barracks":    barracks,
		"Hospital":    hospital,
		"Seedy Bar":   seedyBar,
		"Local Shop":  localShop,
		"Town Square": townSquare,
	}
}

// LinkRooms is a utility function to connect two rooms.
// func LinkRooms(roomA, roomB *Room, direction string) {
// 	switch direction {
// 	case "north":
// 		roomA.N = roomB
// 		roomB.S = roomA
// 	case "south":
// 		roomA.S = roomB
// 		roomB.N = roomA
// 	case "east":
// 		roomA.E = roomB
// 		roomB.W = roomA
// 	case "west":
// 		roomA.W = roomB
// 		roomB.E = roomA
// 	case "enter":
// 		roomA.Enter = roomB
// 		roomB.Exit = roomA
// 	case "exit":
// 		roomA.Exit = roomB
// 		roomB.Enter = roomA
// 	// Add other directions if needed
// 	}
// }

// Description generates a description for the room, including its location in the universe.
func (r *Room) Description(galaxyName, planetName, cityName string) string {
	description := fmt.Sprintf("Galaxy: %s\nPlanet: %s\nCity: %s\nLocation: %s\n\n%s", galaxyName, planetName, cityName, r.Name, r.DescriptionText)

	// Add information about available exits.
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
