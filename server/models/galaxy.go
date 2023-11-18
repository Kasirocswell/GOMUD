// galaxy.go
package models

// Galaxy structure
type Galaxy struct {
	ID        string
	Name      string
	PlanetIDs []string // Array of planet IDs
	Planets   map[string]*Planet
}

// func createGalaxy(name string, planets map[string]*Planet) *Galaxy {
// 	return &Galaxy{
// 		Name:    name,
// 		Planets: planets,
// 	}
// }
