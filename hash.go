package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func calculateFileHash(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:]), nil
}

func handleHash(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: sentinel hash <filename>")
		return
	}

	filename := args[2]

	hashString, err := calculateFileHash(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File:", filename)
	fmt.Println("SHA-256:", hashString)
}