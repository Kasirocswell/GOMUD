package models

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
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
	Soldier  Class = "Soldier"
	Medic    Class = "Medic"
	Pilot    Class = "Pilot"
	Engineer Class = "Engineer"
	Merchant Class = "Merchant"
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

type Effects struct {
	Healthy  string
	Hungry   string
	Poisoned string
}

type User struct {
	Conn           net.Conn
	Name           string
	Race           Race
	Class          Class
	Attributes     Attributes
	Room           *Room
	City           *City
	Planet         *Planet
	Galaxy         *Galaxy
	Reader         *bufio.Reader
	Writer         *bufio.Writer
	State          GameState
	Rolls          int
	Health         int
	MaxHealth      int
	Level          int
	ClassLevel     int
	XP             int
	Inventory      []Item
	Equipment      Equipment
	StatusEffects  []Effects
	CurrentWeight  int            // Total weight of items currently in inventory
	MaxCarryWeight int            // Maximum weight the user can carry
	Skills         map[string]int // Skills with their proficiency levels
	Energy         int            // For actions or abilities
	MaxEnergy      int            // Maximum energy capacity
	Credits        int            // Currency for transactions
	Reputation     map[string]int // Reputation with different factions or groups
	Quests         []Quest        // Active quests
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
	case Merchant:
		u.Class = Merchant
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
		u.setInitialLocation(universe, "cybercluster", "anterra", "neo-tokyo", "central-neotokyo")
	case Medic:
		u.setInitialLocation(universe, "exodus-12", "the-nexus", "solara-crest", "downtown-solara-crest")
	case Pilot:
		u.setInitialLocation(universe, "cybercluster", "silosphere", "del-hi", "uptown-del-hi")
	case Engineer:
		u.setInitialLocation(universe, "prometheous", "echo-prime", "luminon", "central-west-luminon")
	case Merchant:
		u.setInitialLocation(universe, "exodus-12", "haven", "zephyr", "central-zephyr")
	default:
		// Default spawn location
		u.setInitialLocation(universe, "exodus-12", "the-nexus", "nebulon", "downtown-nebulon")
	}
}

func (u *User) setInitialLocation(universe *Universe, galaxyID, planetID, cityID, roomID string) {
	fmt.Printf("Setting location: GalaxyID: %s, PlanetID: %s, CityID: %s, RoomID: %s\n", galaxyID, planetID, cityID, roomID)

	if galaxy := universe.GetGalaxy(galaxyID); galaxy != nil {
		u.Galaxy = galaxy
		fmt.Printf("Found Galaxy: %s\n", galaxy.Name)

		if planet, exists := galaxy.Planets[planetID]; exists {
			u.Planet = planet
			fmt.Printf("Found Planet: %s\n", planet.Name)

			if city, exists := planet.Cities[cityID]; exists {
				u.City = city
				fmt.Printf("Found City: %s\n", city.Name)

				if room, exists := city.Rooms[roomID]; exists {
					u.Room = room
					fmt.Printf("Found Room: %s\n", room.Name)
				} else {
					fmt.Println("Room not found.")
				}
			} else {
				fmt.Println("City not found.")
			}
		} else {
			fmt.Println("Planet not found.")
		}
	} else {
		fmt.Println("Galaxy not found.")
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

//respawn logic

func RunRespawnLoop() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		RespawnDeadEnemies()
	}
}

func RespawnDeadEnemies() {
	currentTime := time.Now()

	// Use a slice to keep track of enemies that need to be removed from DeadEnemies
	var toBeRemoved []int

	for i, enemy := range DeadEnemies {
		if enemy.IsDead && currentTime.Sub(enemy.DeathTime) >= time.Duration(enemy.RespawnTime)*time.Second {
			RespawnEnemy(enemy)
			toBeRemoved = append(toBeRemoved, i)
		}
	}

	// Remove respawned enemies from DeadEnemies list
	for _, index := range toBeRemoved {
		DeadEnemies = append(DeadEnemies[:index], DeadEnemies[index+1:]...)
	}
}

func RespawnEnemy(enemy *Enemy) {
	// Reset enemy properties for respawn (health, position, etc.)
	enemy.IsDead = false
	enemy.Health = enemy.MaxHealth
	// Reset other necessary properties and place the enemy back in the game
}

// AddEnemyToDeadList adds a defeated enemy to the DeadEnemies list
func AddEnemyToDeadList(enemy *Enemy) {
	enemy.IsDead = true
	enemy.DeathTime = time.Now()
	DeadEnemies = append(DeadEnemies, enemy)
}
