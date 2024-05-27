package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer ln.Close() // Close the listener when main function exits
	fmt.Println("Server is listening on port 8080...")

	// Accept incoming connections and handle them
	for {
		conn, err := ln.Accept() // Accept new connections
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // Close the connection when handleConnection exits

	// Read incoming data continuously
	buf := make([]byte, 1024) // Buffer to hold incoming data
	for {
		n, err := conn.Read(buf) // Read data from the connection
		if err != nil {
			if err.Error() != "EOF" { // Check for EOF error
				fmt.Println("Error reading from connection:", err)
			}
			return // Return if there is an error
		}

		// Print the incoming data
		if n > 0 {
			fmt.Printf("Received: %s\n", string(buf[:n]))
		}
	}
}
