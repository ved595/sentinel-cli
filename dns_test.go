package main

import "testing"

func TestLookupHost(t *testing.T) {
	addresses, err := lookupHost("localhost")

	if err != nil {
		t.Fatal("DNS lookup failed:", err)
	}

	if len(addresses) == 0 {
		t.Fatal("Expected at least one IP address")
	}
}