package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func checkPort(host string, port string) bool {
	address := net.JoinHostPort(host, port)

	connection, err := net.DialTimeout(
		"tcp",
		address,
		3*time.Second,
	)

	if err != nil {
		return false
	}

	defer connection.Close()

	return true
}

func handlePort(args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: sentinel port <host> <port>")
		return
	}

	host := args[2]
	port := args[3]

	if host == "" {
		fmt.Println("Error: host cannot be empty")
		return
	}

	portNumber, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Error: port must be a number")
		return
	}

	if portNumber < 1 || portNumber > 65535 {
		fmt.Println("Error: port must be between 1 and 65535")
		return
	}

	fmt.Println("Checking:", net.JoinHostPort(host, port))

	isOpen := checkPort(host, port)

	if isOpen {
		fmt.Println("Status: OPEN")
	} else {
		fmt.Println("Status: CLOSED or unreachable")
	}
}