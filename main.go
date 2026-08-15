package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sentinel <command>")
		return
	}

	command := os.Args[1]

	if command == "hash" {
	if len(os.Args) < 3 {
		fmt.Println("Usage: sentinel hash <filename>")
		return
	}

	filename := os.Args[2]

data, err := os.ReadFile(filename)
if err != nil {
	fmt.Println("Error reading file:", err)
	return
}

hash := sha256.Sum256(data)

fmt.Println("File:", filename)
fmt.Println("SHA-256:", hex.EncodeToString(hash[:]))

		return
	}

	fmt.Println("Unknown command:", command)
}