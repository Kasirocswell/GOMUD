package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// LoadGalaxies loads galaxy data from a JSON file.
func LoadGalaxies(filename string) ([]Galaxy, error) {
	var galaxies []Galaxy

	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read file content
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse JSON into the slice of Galaxy structs
	err = json.Unmarshal(byteValue, &galaxies)
	if err != nil {
		return nil, err
	}

	return galaxies, nil
}

// LoadPlanets loads planet data from a JSON file.
func LoadPlanets(filename string) ([]Planet, error) {
	var planets []Planet

	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Decode the JSON data into the slice of Planet structs
	err = json.NewDecoder(file).Decode(&planets)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return planets, nil
}

// LoadCities loads city data from a JSON file.
func LoadCities(filename string) ([]City, error) {
	var cities []City

	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Decode the JSON data into the slice of City structs
	err = json.NewDecoder(file).Decode(&cities)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return cities, nil
}

func LoadRoomsFromFolder(folderPath string) (map[string][]Room, error) {
	roomsByCity := make(map[string][]Room)

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			var rooms []Room
			fileData, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if err := json.Unmarshal(fileData, &rooms); err != nil {
				return err
			}
			city := filepath.Base(path)
			city = city[:len(city)-len(filepath.Ext(city))] // Strip file extension
			roomsByCity[city] = rooms
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return roomsByCity, nil
}
