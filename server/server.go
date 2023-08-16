package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type User struct {
	conn   net.Conn
	name   string
	room   *Room
	reader *bufio.Reader
	writer *bufio.Writer
}

type Room struct {
	description string
	users       map[*User]bool
}

type World struct {
	rooms map[string]*Room
}

func (u *User) readInput() string {
	input, _ := u.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (u *User) handleCommand(command string) {
	// Handle different commands
	if command == "look" {
		u.writer.WriteString(u.room.description + "\n")
		u.writer.Flush()
	} else {
		u.writer.WriteString("Unknown command.\n")
		u.writer.Flush()
	}
}

func handleConnection(conn net.Conn, world *World) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	user := &User{
		conn:   conn,
		room:   world.rooms["start"],
		reader: r,
		writer: w,
	}

	// Assume we just drop the user into the starting room for simplicity
	for {
		input := user.readInput()
		user.handleCommand(input)
	}
}

func main() {
	// Create world and rooms
	world := &World{
		rooms: make(map[string]*Room),
	}
	world.rooms["start"] = &Room{
		description: "You are in a starting room.",
		users:       make(map[*User]bool),
	}

	// Start server
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("MUD server started on localhost:4000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn, world)
	}
}
