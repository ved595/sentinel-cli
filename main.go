package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "hash":
		handleHash(os.Args)

	case "baseline":
		handleBaseline(os.Args)

	case "check":
		handleCheck(os.Args)

	case "dns":
		handleDNS(os.Args)

	case "port":
		handlePort(os.Args)

	default:
		fmt.Println("Unknown command:", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Sentinel CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  sentinel hash <filename>")
	fmt.Println("  sentinel baseline <filename>")
	fmt.Println("  sentinel check <filename>")
	fmt.Println("  sentinel dns <domain>")
	fmt.Println("  sentinel port <host> <port>")
}