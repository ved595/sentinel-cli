package main

import (
	"fmt"
	"net"
	"strings"
)

func lookupHost(domain string) ([]string, error) {
	return net.LookupHost(domain)
}

func handleDNS(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: sentinel dns <domain>")
		return
	}

	domain := strings.TrimSpace(args[2])

	if domain == "" {
		fmt.Println("Error: domain cannot be empty")
		return
	}

	ipAddresses, err := lookupHost(domain)
	if err != nil {
		fmt.Println("DNS lookup failed:", err)
		return
	}

	fmt.Println("Domain:", domain)
	fmt.Println("IP addresses:")

	for _, ip := range ipAddresses {
		fmt.Println("-", ip)
	}
}