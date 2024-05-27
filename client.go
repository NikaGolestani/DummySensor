package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Changed to UnixNano for nanosecond precision

	reader := bufio.NewReader(os.Stdin)

	// Prompt user for the interval between sending data
	fmt.Print("Enter interval between sending data in nanoseconds: ")
	intervalInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Convert input string to integer for interval
	interval, err := strconv.Atoi(strings.TrimSpace(intervalInput))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer for interval.")
		return
	}

	// Prompt user for the minimum value for x
	fmt.Print("Enter minimum value for x: ")
	minXInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Convert input string to integer for minimum value of x
	minX, err := strconv.Atoi(strings.TrimSpace(minXInput))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer for minimum value of x.")
		return
	}

	// Prompt user for the maximum value for x
	fmt.Print("Enter maximum value for x: ")
	maxXInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Convert input string to integer for maximum value of x
	maxX, err := strconv.Atoi(strings.TrimSpace(maxXInput))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer for maximum value of x.")
		return
	}

	// Ensure minX is less than maxX
	if minX >= maxX {
		fmt.Println("Minimum value for x must be less than maximum value.")
		return
	}

	// Prompt user for the minimum value for y
	fmt.Print("Enter minimum value for y: ")
	minYInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Convert input string to integer for minimum value of y
	minY, err := strconv.Atoi(strings.TrimSpace(minYInput))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer for minimum value of y.")
		return
	}

	// Prompt user for the maximum value for y
	fmt.Print("Enter maximum value for y: ")
	maxYInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Convert input string to integer for maximum value of y
	maxY, err := strconv.Atoi(strings.TrimSpace(maxYInput))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer for maximum value of y.")
		return
	}

	// Ensure minY is less than maxY
	if minY >= maxY {
		fmt.Println("Minimum value for y must be less than maximum value.")
		return
	}

	for {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			time.Sleep(time.Second * 5) // Wait before retrying
			continue
		}

		for {
			// Generate random values for x and y within specified range
			x := rand.Intn(maxX-minX+1) + minX
			y := rand.Intn(maxY-minY+1) + minY

			// Construct the message to be sent to the server
			message := strconv.Itoa(x) + ":" + strconv.Itoa(y) + "\n"

			// Send the message to the server
			_, err = conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Error writing to connection:", err)
				conn.Close()
				break // Exit the inner loop to retry the connection
			}

			// Wait for the specified interval before sending the next message
			time.Sleep(time.Duration(interval) * time.Nanosecond)
		}

		// Ensure the connection is closed before retrying
		conn.Close()
	}
}
