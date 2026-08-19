package main

import (
	"os"
	"testing"
)

func TestCalculateFileHash(t *testing.T) {
	filename := "test-hash.txt"
	content := []byte("hello sentinel")

	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		t.Fatal("Could not create test file:", err)
	}

	defer os.Remove(filename)

	hash, err := calculateFileHash(filename)
	if err != nil {
		t.Fatal("Hash calculation failed:", err)
	}

	expected := "8c44dc2aab50e05dd4e8090788add36284ea7c17513cd5d809ff9ca9152fa4c2"

	if hash != expected {
		t.Errorf("Expected %s, got %s", expected, hash)
	}
}