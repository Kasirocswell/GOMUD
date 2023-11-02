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
	Draconian Race = "Draconian"
	Cyborg    Race = "Cyborg"
	Uthalu    Race = "Uthalu"
	Drogan    Race = "Drogan"
)

type Class string

const (
	Soldier      Class = "Soldier"
	Medic        Class = "Medic"
	Pilot        Class = "Pilot"
	Engineer     Class = "Engineer"
	Entrepreneur Class = "Entrepreneur"
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
		u.Galaxy = universe.GetGalaxy("CyberCluster")
		u.Planet = universe.GetPlanet("CyberCluster", "Techterra")
		u.City = universe.GetCity("CyberCluster", "Techterra", "NeoTokyo")
		u.Room = universe.GetRoom("CyberCluster", "Techterra", "NeoTokyo", "Barracks")
	case Medic:
		u.Galaxy = universe.GetGalaxy("DigitalDomain")
		u.Planet = universe.GetPlanet("DigitalDomain", "SiliconSphere")
		u.City = universe.GetCity("DigitalDomain", "SiliconSphere", "SiliconParis")
		u.Room = universe.GetRoom("DigitalDomain", "SiliconSphere", "SiliconParis", "Hospital")
	case Pilot:
		u.Galaxy = universe.GetGalaxy("NetNebula")
		u.Planet = universe.GetPlanet("NetNebula", "NetNeptune")
		u.City = universe.GetCity("NetNebula", "NetNeptune", "NeptuneCity1")
		u.Room = universe.GetRoom("NetNebula", "NetNeptune", "NeptuneCity1", "Seedy Bar")
	case Engineer:
		u.Galaxy = universe.GetGalaxy("TechTwilight")
		u.Planet = universe.GetPlanet("TechTwilight", "Digitalis")
		u.City = universe.GetCity("TechTwilight", "Digitalis", "DigitalisCity1")
		u.Room = universe.GetRoom("TechTwilight", "Digitalis", "DigitalisCity1", "Local Shop")
	case Entrepreneur:
		u.Galaxy = universe.GetGalaxy("CyberCluster")
		u.Planet = universe.GetPlanet("CyberCluster", "Techterra")
		u.City = universe.GetCity("CyberCluster", "Techterra", "CyberLisbon")
		u.Room = universe.GetRoom("CyberCluster", "Techterra", "CyberLisbon", "Town Square")
	default:
		// Default spawn location if none of the above classes match.
		u.Galaxy = universe.GetGalaxy("CyberCluster")
		u.Planet = universe.GetPlanet("CyberCluster", "Techterra")
		u.City = universe.GetCity("CyberCluster", "Techterra", "NeoTokyo")
		u.Room = universe.GetRoom("CyberCluster", "Techterra", "NeoTokyo", "Town Square")
	}

	// Check if any of the locations are nil
	if u.Galaxy == nil {
		fmt.Println("DEBUG: Galaxy is nil!")
	}
	if u.Planet == nil {
		fmt.Println("DEBUG: Planet is nil!")
	}
	if u.City == nil {
		fmt.Println("DEBUG: City is nil!")
	}
	if u.Room == nil {
		fmt.Println("DEBUG: Room is nil!")
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
