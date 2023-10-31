// galaxy.go
package models

// Galaxy structure
type Galaxy struct {
	Name    string
	Planets map[string]*Planet
}

func createGalaxy(name string, planets map[string]*Planet) *Galaxy {
	return &Galaxy{
		Name:    name,
		Planets: planets,
	}
}
