package main

import (
	"fmt"
	"net"
)

func handleDNS(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: sentinel dns <domain>")
		return
	}

	domain := args[2]

	ipAddresses, err := net.LookupHost(domain)
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
