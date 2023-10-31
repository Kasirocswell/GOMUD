package models

import (
	"fmt"
)

// Room structure
type Room struct {
	Name                string
	DescriptionTemplate string
	Users               map[*User]bool
	Exits               map[string]*Room
}

// Description generates a description for the room based on its template.
func (r *Room) Description(cityName, planetName string) string {
	return fmt.Sprintf(r.DescriptionTemplate, cityName, planetName)
}

// CreateCommonRoom creates a common room with a given name and description template.
func CreateCommonRoom(name, descriptionTemplate string) *Room {
	return &Room{
		Name:                name,
		DescriptionTemplate: descriptionTemplate,
		Users:               make(map[*User]bool),
		Exits:               make(map[string]*Room),
	}
}

// CreateUniqueRoom creates a unique room with a given name and description.
func CreateUniqueRoom(name, description string) *Room {
	return &Room{
		Name:                name,
		DescriptionTemplate: description,
		Users:               make(map[*User]bool),
		Exits:               make(map[string]*Room),
	}
}
