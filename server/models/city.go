package models

// City structure
type City struct {
	ID       string
	PlanetID string // Link to the planet
	Name     string
	Rooms    map[string]*Room
}
