package models

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
)

// Enums and Data Structures
type GameState int

const (
	CharacterCreation GameState = iota
	InGame
)

type Race string

const (
	Human     Race = "Human"
	Draconian      = "Draconian"
	Cyborg         = "Cyborg"
	Uthalu         = "Uthalu"
	Drogan         = "Drogan"
)

type Class string

const (
	Soldier      Class = "Soldier"
	Medic              = "Medic"
	Pilot              = "Pilot"
	Engineer           = "Engineer"
	Entrepreneur       = "Entrepreneur"
)

type Attributes struct {
	Strength     int
	Endurance    int
	Speed        int
	Agility      int
	Intelligence int
	Charisma     int
	Luck         int
	Perception   int
	Wisdom       int
}

type User struct {
	Conn       net.Conn
	Name       string
	Race       Race
	Class      Class
	Attributes Attributes
	Room       *Room
	City       *City
	Planet     *Planet
	Galaxy     *Galaxy
	Reader     *bufio.Reader
	Writer     *bufio.Writer
	State      GameState
	Rolls      int
}

func NewUser(conn net.Conn, r *bufio.Reader, w *bufio.Writer) *User {
	return &User{
		Conn:   conn,
		Reader: r,
		Writer: w,
		State:  CharacterCreation,
		Rolls:  1, // Starting with the initial roll
	}
}

// SetRace sets the user's race and handles invalid input.
// SetRace sets the user's race and handles invalid input.
func (u *User) SetRace(raceChoice string) error {
	switch Race(raceChoice) {
	case Human:
		u.Race = Human
	case Draconian:
		u.Race = Draconian
	case Cyborg:
		u.Race = Cyborg
	case Uthalu:
		u.Race = Uthalu
	case Drogan:
		u.Race = Drogan
	default:
		return fmt.Errorf("invalid race selection: %s; please choose from Human, Draconian, Cyborg, Uthalu, or Drogan", raceChoice)
	}
	return nil
}

// SetClass sets the user's class and handles invalid input.

func (u *User) SetClass(classChoice string) error {
	switch Class(classChoice) {
	case Soldier:
		u.Class = Soldier
	case Medic:
		u.Class = Medic
	case Pilot:
		u.Class = Pilot
	case Engineer:
		u.Class = Engineer
	case Entrepreneur:
		u.Class = Entrepreneur
	default:
		return fmt.Errorf("invalid class selection: %s; please choose from Soldier, Medic, Pilot, Engineer, or Entrepreneur", classChoice)
	}
	return nil
}

func (u *User) RollAttributes() {
	// Seed the random number generator to ensure different results each time
	// rand.Seed(time.Now().UnixNano())

	// Assign random values to each attribute. For this example, I'm assuming a range of 1-10 for each attribute.
	u.Attributes.Strength = rand.Intn(10) + 1
	u.Attributes.Endurance = rand.Intn(10) + 1
	u.Attributes.Speed = rand.Intn(10) + 1
	u.Attributes.Agility = rand.Intn(10) + 1
	u.Attributes.Intelligence = rand.Intn(10) + 1
	u.Attributes.Charisma = rand.Intn(10) + 1
	u.Attributes.Luck = rand.Intn(10) + 1
	u.Attributes.Perception = rand.Intn(10) + 1
	u.Attributes.Wisdom = rand.Intn(10) + 1
}

// DisplayAttributes returns a string representation of the user's attributes.
func (u *User) DisplayAttributes() string {
	return fmt.Sprintf(
		"Attributes:\n"+
			"Strength: %d\n"+
			"Endurance: %d\n"+
			"Speed: %d\n"+
			"Agility: %d\n"+
			"Intelligence: %d\n"+
			"Charisma: %d\n"+
			"Luck: %d\n"+
			"Perception: %d\n"+
			"Wisdom: %d\n",
		u.Attributes.Strength,
		u.Attributes.Endurance,
		u.Attributes.Speed,
		u.Attributes.Agility,
		u.Attributes.Intelligence,
		u.Attributes.Charisma,
		u.Attributes.Luck,
		u.Attributes.Perception,
		u.Attributes.Wisdom,
	)
}

// SpawnInUniverse sets the user's initial location based on their class.
func (u *User) SpawnInUniverse(universe *Universe) {
	switch u.Class {
	case Soldier:
		u.Galaxy = universe.GetGalaxy("Milky Way")
		u.Planet = universe.GetPlanet("Milky Way", "Earth")
		u.City = universe.GetCity("Milky Way", "Earth", "Military Base")
		u.Room = universe.GetRoom("Milky Way", "Earth", "Military Base", "Barracks")
	case Medic:
		u.Galaxy = universe.GetGalaxy("Andromeda")
		u.Planet = universe.GetPlanet("Andromeda", "Healix")
		u.City = universe.GetCity("Andromeda", "Healix", "Medic Central")
		u.Room = universe.GetRoom("Andromeda", "Healix", "Medic Central", "Infirmary")
	case Pilot:
		u.Galaxy = universe.GetGalaxy("Nebula X")
		u.Planet = universe.GetPlanet("Nebula X", "Skytopia")
		u.City = universe.GetCity("Nebula X", "Skytopia", "Pilot's Haven")
		u.Room = universe.GetRoom("Nebula X", "Skytopia", "Pilot's Haven", "Hangar")
	case Engineer:
		u.Galaxy = universe.GetGalaxy("Techno Realm")
		u.Planet = universe.GetPlanet("Techno Realm", "Gearland")
		u.City = universe.GetCity("Techno Realm", "Gearland", "Engineer's Enclave")
		u.Room = universe.GetRoom("Techno Realm", "Gearland", "Engineer's Enclave", "Workshop")
	case Entrepreneur:
		u.Galaxy = universe.GetGalaxy("Golden Cluster")
		u.Planet = universe.GetPlanet("Golden Cluster", "Marketron")
		u.City = universe.GetCity("Golden Cluster", "Marketron", "Trade Town")
		u.Room = universe.GetRoom("Golden Cluster", "Marketron", "Trade Town", "Marketplace")
	default:
		// Default spawn location if none of the above classes match.
		u.Galaxy = universe.GetGalaxy("Milky Way")
		u.Planet = universe.GetPlanet("Milky Way", "Earth")
		u.City = universe.GetCity("Milky Way", "Earth", "Downtown")
		u.Room = universe.GetRoom("Milky Way", "Earth", "Downtown", "Central Square")
	}
}

// User Input
func (u *User) ReadInput() (string, error) {
	input, err := u.Reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// User Commands
func (u *User) HandleCommand(command string) {
	switch command {
	case "look":
		u.Writer.WriteString(u.Room.Description(u.City.Name, u.Planet.Name) + "\n")
	case "n", "north":
		u.move("north")
	case "s", "south":
		u.move("south")
	case "e", "east":
		u.move("east")
	case "w", "west":
		u.move("west")
	case "enter":
		u.move("enter_building") // or whatever key you use for entering
	case "exit":
		u.move("exit") // or whatever key you use for exiting
	default:
		u.Writer.WriteString("Unknown command.\n")
	}
	u.Writer.Flush()
}

func (u *User) move(direction string) {
	if exit, ok := u.Room.Exits[direction]; ok {
		u.Room = exit
		u.Writer.WriteString("You move " + direction + " to " + u.Room.Name + ".\n")
		u.Writer.WriteString(u.Room.Description(u.City.Name, u.Planet.Name) + "\n")
	} else {
		u.Writer.WriteString("There's no exit in that direction.\n")
	}
	u.Writer.Flush()
}
