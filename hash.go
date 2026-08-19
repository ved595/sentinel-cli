package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func handleHash(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: sentinel hash <filename>")
		return
	}

	filename := args[2]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	hash := sha256.Sum256(data)
	hashString := hex.EncodeToString(hash[:])

	fmt.Println("File:", filename)
	fmt.Println("SHA-256:", hashString)
}