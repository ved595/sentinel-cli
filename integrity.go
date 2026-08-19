package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const baselineFile = "sentinel-baseline.txt"

func handleBaseline(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: sentinel baseline <filename>")
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

	baseline := filename + ":" + hashString

	err = os.WriteFile(
		baselineFile,
		[]byte(baseline),
		0644,
	)
	if err != nil {
		fmt.Println("Error saving baseline:", err)
		return
	}

	fmt.Println("Baseline created for:", filename)
}

func handleCheck(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: sentinel check <filename>")
		return
	}

	filename := args[2]

	baselineData, err := os.ReadFile(baselineFile)
	if err != nil {
		fmt.Println("Error reading baseline:", err)
		return
	}

	parts := strings.SplitN(string(baselineData), ":", 2)
	if len(parts) != 2 {
		fmt.Println("Invalid baseline file")
		return
	}

	savedFilename := parts[0]
	savedHash := parts[1]

	if savedFilename != filename {
		fmt.Println("No baseline found for:", filename)
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	currentHash := sha256.Sum256(data)
	currentHashString := hex.EncodeToString(currentHash[:])

	if currentHashString == savedHash {
		fmt.Println("File integrity verified:", filename)
	} else {
		fmt.Println("WARNING: File has been modified:", filename)
	}
}