package main

import (
	"os"
	"testing"
)

func TestFileIntegrity(t *testing.T) {
	filename := "integrity-test.txt"
	originalContent := []byte("trusted content")

	err := os.WriteFile(filename, originalContent, 0644)
	if err != nil {
		t.Fatal("Could not create test file:", err)
	}

	defer os.Remove(filename)
	defer os.Remove(baselineFile)

	err = createBaseline(filename)
	if err != nil {
		t.Fatal("Could not create baseline:", err)
	}

	matches, err := checkFileIntegrity(filename)
	if err != nil {
		t.Fatal("Integrity check failed:", err)
	}

	if !matches {
		t.Fatal("Expected file to match baseline")
	}

	modifiedContent := []byte("modified content")

	err = os.WriteFile(filename, modifiedContent, 0644)
	if err != nil {
		t.Fatal("Could not modify test file:", err)
	}

	matches, err = checkFileIntegrity(filename)
	if err != nil {
		t.Fatal("Integrity check failed:", err)
	}

	if matches {
		t.Fatal("Expected modified file to fail integrity check")
	}
}