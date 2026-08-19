# Sentinel CLI

Sentinel CLI is a command-line cybersecurity and networking toolkit built with Go.

The project explores fundamental cybersecurity concepts through practical utilities for file integrity monitoring, DNS resolution, and TCP connectivity testing.

## Features

### SHA-256 File Hashing

Generate the SHA-256 hash of a file.

```bash
go run . hash <filename>
```

Example:

```bash
go run . hash test.txt
```

Output:

```text
File: test.txt
SHA-256: <generated SHA-256 hash>
```

### File Integrity Baselines

Create a trusted SHA-256 baseline for a file.

```bash
go run . baseline <filename>
```

Example:

```bash
go run . baseline test.txt
```

Output:

```text
Baseline created for: test.txt
```

Sentinel stores the trusted hash locally and can later compare the file against that baseline.

### File Integrity Verification

Check whether a monitored file has changed:

```bash
go run . check <filename>
```

If the file matches its baseline:

```text
File integrity verified: test.txt
```

If the contents have changed:

```text
WARNING: File has been modified: test.txt
```

This demonstrates the basic concept behind file integrity monitoring: detecting unexpected changes by comparing cryptographic hashes.

### DNS Lookup

Resolve a domain name to its associated IP addresses:

```bash
go run . dns <domain>
```

Example:

```bash
go run . dns example.com
```

Example output:

```text
Domain: example.com
IP addresses:
- <IPv4 address>
- <IPv6 address>
```

Results may contain both IPv4 and IPv6 addresses.

### TCP Port Checking

Test whether a specific TCP port accepts a connection:

```bash
go run . port <host> <port>
```

Example:

```bash
go run . port example.com 443
```

Example output:

```text
Checking: example.com:443
Status: OPEN
```

If Sentinel cannot establish a connection:

```text
Status: CLOSED or unreachable
```

Port numbers are validated and must be between `1` and `65535`.

## CLI Commands

```text
sentinel hash <filename>
sentinel baseline <filename>
sentinel check <filename>
sentinel dns <domain>
sentinel port <host> <port>
```

Running Sentinel without a command displays the available commands.

## Project Structure

```text
sentinel-cli/
├── main.go
├── hash.go
├── integrity.go
├── dns.go
├── port.go
├── hash_test.go
├── integrity_test.go
├── dns_test.go
├── port_test.go
├── go.mod
├── .gitignore
└── README.md
```

The project separates command handling from the underlying functionality so individual components can be tested independently.

## Automated Testing

Sentinel includes automated tests for its core functionality.

Run the complete test suite with:

```bash
go test
```

The tests cover:

- SHA-256 hash calculation
- Missing-file error handling
- File integrity verification
- File modification detection
- DNS resolution
- TCP connection checking

Networking tests use local resources where possible so the test suite does not unnecessarily depend on external services.

## Building Sentinel

Build the executable with:

```bash
go build -o sentinel
```

Then run Sentinel directly:

```bash
./sentinel
```

For example:

```bash
./sentinel dns example.com
```

or:

```bash
./sentinel port example.com 443
```

## Security Concepts Demonstrated

This project demonstrates several foundational cybersecurity and networking concepts:

- Cryptographic hashing with SHA-256
- File integrity monitoring
- Change detection
- DNS resolution
- IPv4 and IPv6
- TCP connections
- Network ports
- Connection timeouts
- Input validation
- Error handling
- Automated testing

## Responsible Use

Sentinel's networking functionality is intended for education, network troubleshooting, and testing systems you own or have explicit permission to test.

## Technologies

- Go
- Git
- GitHub
- Go standard library

## Project Status

Sentinel CLI v1 is feature complete.

Future versions may expand the toolkit with additional security and networking functionality.