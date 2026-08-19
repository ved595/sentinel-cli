package main

import (
	"fmt"
	"net"
	"time"
)

func handlePort(args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: sentinel port <host> <port>")
		return
	}

	host := args[2]
	port := args[3]

	address := net.JoinHostPort(host, port)

	fmt.Println("Checking:", address)

	connection, err := net.DialTimeout(
		"tcp",
		address,
		3*time.Second,
	)

	if err != nil {
		fmt.Println("Status: CLOSED or unreachable")
		return
	}

	defer connection.Close()

	fmt.Println("Status: OPEN")
}
