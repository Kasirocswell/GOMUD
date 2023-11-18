// handlers/connection_handlers.go

package handlers

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/kasirocswell/gomudserver/models"
)

func HandleNewConnection(conn net.Conn, universe *models.Universe) {
	defer conn.Close()

	user := &models.User{
		Conn:   conn,
		Reader: bufio.NewReader(conn),
		Writer: bufio.NewWriter(conn),
		State:  models.CharacterCreation,
	}

	user.Writer.WriteString("Welcome to Hacker-Realm!\n")
	user.Writer.WriteString("Please Enter Your Character's Name:\n")
	user.Writer.Flush()

	for {
		input, err := user.ReadInput()
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		switch user.State {
		case models.CharacterCreation:
			if user.Name == "" {
				user.Name = input
				user.Writer.WriteString("Select your race (Human, Draconian, Cyborg, Uthalu, Drogan): ")
				user.Writer.Flush()
			} else if user.Race == "" {
				err := user.SetRace(input)
				if err != nil {
					user.Writer.WriteString(err.Error() + "\n")
					user.Writer.Flush()
					continue
				}
				user.Writer.WriteString("Select your class (Soldier, Medic, Pilot, Engineer, Entrepreneur): ")
				user.Writer.Flush()
			} else if user.Class == "" {
				err := user.SetClass(input)
				if err != nil {
					user.Writer.WriteString(err.Error() + "\n")
					user.Writer.Flush()
					continue
				}
				user.RollAttributes()
				user.Writer.WriteString(user.DisplayAttributes())
				user.Writer.WriteString("You can choose to 'keep' these attributes or 'reroll'. You have a total of 3 rolls including the initial one.\n")
				user.Writer.Flush()
			} else if user.Rolls < 3 {
				if input == "keep" {
					user.State = models.InGame
					// Debug print statement to check the user's game state and other attributes
					fmt.Printf("DEBUG: User State: %v, Name: %s, Race: %s, Class: %s, Attributes: %+v\n", user.State, user.Name, user.Race, user.Class, user.Attributes)
					user.SpawnInUniverse(universe)
					roomName := ""
					cityName := ""
					planetName := ""
					galaxyName := ""

					if user.Room != nil {
						roomName = user.Room.Name
					}
					if user.City != nil {
						cityName = user.City.Name
					}
					if user.Planet != nil {
						planetName = user.Planet.Name
					}
					if user.Galaxy != nil {
						galaxyName = user.Galaxy.Name
					}
					user.Writer.WriteString(fmt.Sprintf("You have been spawned in %s, %s on planet %s in the %s galaxy.\n", roomName, cityName, planetName, galaxyName))
					fmt.Println("DEBUG: Room:", user.Room)
					fmt.Println("DEBUG: City:", user.City)
					fmt.Println("DEBUG: Planet:", user.Planet)
					fmt.Println("DEBUG: Galaxy:", user.Galaxy)

					// user.Writer.WriteString(fmt.Sprintf("You have been spawned in %s, %s on planet %s in the %s galaxy.\n", user.Room.Name, user.City.Name, user.Planet.Name, user.Galaxy.Name))
					user.Writer.Flush()
				} else if input == "reroll" {
					user.RollAttributes()
					user.Rolls++
					user.Writer.WriteString(user.DisplayAttributes())
					user.Writer.WriteString("You can choose to 'keep' these attributes or 'reroll'. You have a total of 3 rolls including the initial one.\n")
					user.Writer.Flush()
				} else {
					user.Writer.WriteString("Invalid choice. Please choose 'keep' or 'reroll'.\n")
					user.Writer.Flush()
				}
			} else {
				user.State = models.InGame
				// Extract command and arguments
				command, args := parseCommand(input)
				HandleCommand(user, command, universe, args...)
				user.SpawnInUniverse(universe)
				user.Writer.WriteString(fmt.Sprintf("You have been spawned in %s, %s on planet %s in the %s galaxy.\n", user.Room.Name, user.City.Name, user.Planet.Name, user.Galaxy.Name))
				user.Writer.Flush()
			}
		case models.InGame:
			// Handle in-game commands using the abstracted function
			command, args := parseCommand(input)
			HandleCommand(user, command, universe, args...)

		}
	}
}

// parseCommand splits a command string into the command itself and additional arguments.
func parseCommand(input string) (string, []string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", nil
	}

	command := parts[0]
	args := []string{}
	if len(parts) > 1 {
		args = parts[1:]
	}

	return command, args
}

// welcomeUser sends a welcome message to the user.
// func welcomeUser(user *models.User) {
// 	user.Writer.WriteString("Welcome to GoMUD!\n")
// 	user.Writer.WriteString("Please enter a command or type 'help' for a list of commands.\n")
// 	user.Writer.Flush()
// }
