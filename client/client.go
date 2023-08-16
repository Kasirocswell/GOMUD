package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:4000")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    go func() {
        scanner := bufio.NewScanner(conn)
        for scanner.Scan() {
            fmt.Println(scanner.Text())
        }
    }()

    clientScanner := bufio.NewScanner(os.Stdin)
    for clientScanner.Scan() {
        fmt.Fprintln(conn, clientScanner.Text())
    }
}
