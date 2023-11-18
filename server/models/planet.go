// planet.go
package models

// Planet structure
type Planet struct {
	ID       string
	GalaxyID string
	Name     string
	Cities   map[string]*City
}

func CreatePlanet(name string, cities map[string]*City) *Planet {
	return &Planet{
		Name:   name,
		Cities: cities,
	}
}
