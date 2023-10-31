// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net"

// 	"github.com/kasirocswell/gomudserver/models" // Replace with the actual path to your models package
// )

// func handleConnection(conn net.Conn, universe *models.Universe) {
// 	defer conn.Close()

// 	r := bufio.NewReader(conn)
// 	w := bufio.NewWriter(conn)

// 	user := models.NewUser(conn, r, w)

// 	w.WriteString("Welcome to Hacker-Realm!\n")
// 	w.WriteString("Please Enter Your Character's Name:\n")
// 	w.Flush()

// 	for {
// 		input, err := user.ReadInput()
// 		if err != nil {
// 			fmt.Println("Error reading input:", err)
// 			return
// 		}

// 		switch user.State {
// 		case models.CharacterCreation:
// 			if user.Name == "" {
// 				user.Name = input
// 				w.WriteString("Select your race (Human, Draconian, Cyborg, Uthalu, Drogan): ")
// 				w.Flush()
// 			} else if user.Race == "" {
// 				user.SetRace(input)
// 				w.WriteString("Select your class (Soldier, Medic, Pilot, Engineer, Entrepreneur): ")
// 				w.Flush()
// 			} else if user.Class == "" {
// 				user.SetClass(input)
// 				user.RollAttributes()
// 				w.WriteString(user.DisplayAttributes())
// 				w.WriteString("You can choose to 'keep' these attributes or 'reroll'. You have a total of 3 rolls including the initial one.\n")
// 				w.Flush()
// 			} else if user.Rolls < 3 {
// 				if input == "keep" {
// 					user.State = models.InGame
// 					user.SpawnInUniverse(universe)
// 					w.WriteString(fmt.Sprintf("You have been spawned in %s, %s on planet %s in the %s galaxy.\n", user.Room.Name, user.City.Name, user.Planet.Name, user.Galaxy.Name))
// 					w.Flush()
// 				} else if input == "reroll" {
// 					user.RollAttributes()
// 					w.WriteString(user.DisplayAttributes())
// 					w.WriteString("You can choose to 'keep' these attributes or 'reroll'. You have a total of 3 rolls including the initial one.\n")
// 					w.Flush()
// 				} else {
// 					w.WriteString("Invalid choice. Please choose 'keep' or 'reroll'.\n")
// 					w.Flush()
// 				}
// 			} else {
// 				user.State = models.InGame
// 				user.SpawnInUniverse(universe)
// 				w.WriteString(fmt.Sprintf("You have been spawned in %s, %s on planet %s in the %s galaxy.\n", user.Room.Name, user.City.Name, user.Planet.Name, user.Galaxy.Name))
// 				w.Flush()
// 			}

// 		case models.InGame:
// 			// Handle in-game commands
// 			user.HandleCommand(input)
// 		}
// 	}
// }

// func main() {
// 	universe := models.CreateUniverse()

// 	listener, err := net.Listen("tcp", "localhost:4000")
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 		return
// 	}
// 	defer listener.Close()

// 	fmt.Println("MUD server started on localhost:4000")

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting connection:", err)
// 			continue
// 		}
// 		go handleConnection(conn, universe)
// 	}
// }
