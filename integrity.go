package main

import (
	"fmt"
	"os"
	"strings"
)

const baselineFile = "sentinel-baseline.txt"

func createBaseline(filename string) error {
	hashString, err := calculateFileHash(filename)
	if err != nil {
		return err
	}

	baseline := filename + ":" + hashString

	err = os.WriteFile(
		baselineFile,
		[]byte(baseline),
		0644,
	)
	if err != nil {
		return err
	}

	return nil
}

func checkFileIntegrity(filename string) (bool, error) {
	baselineData, err := os.ReadFile(baselineFile)
	if err != nil {
		return false, err
	}

	parts := strings.SplitN(string(baselineData), ":", 2)

	if len(parts) != 2 {
		return false, fmt.Errorf("invalid baseline file")
	}

	savedFilename := parts[0]
	savedHash := parts[1]

	if savedFilename != filename {
		return false, fmt.Errorf("no baseline found for %s", filename)
	}

	currentHash, err := calculateFileHash(filename)
	if err != nil {
		return false, err
	}

	return currentHash == savedHash, nil
}

func handleBaseline(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: sentinel baseline <filename>")
		return
	}

	filename := args[2]

	err := createBaseline(filename)
	if err != nil {
		fmt.Println("Error creating baseline:", err)
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

	matches, err := checkFileIntegrity(filename)
	if err != nil {
		fmt.Println("Integrity check failed:", err)
		return
	}

	if matches {
		fmt.Println("File integrity verified:", filename)
	} else {
		fmt.Println("WARNING: File has been modified:", filename)
	}
}