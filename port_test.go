package main

import (
	"fmt"
	"net"
	"testing"
)

func TestCheckPort(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		t.Fatal("Could not create test server:", err)
	}

	defer listener.Close()

	address := listener.Addr().(*net.TCPAddr)

	host := "127.0.0.1"
	port := address.Port

	isOpen := checkPort(host, fmt.Sprintf("%d", port))

	if !isOpen {
		t.Fatal("Expected test port to be open")
	}
}
