# Sentinel CLI

Sentinel CLI is a command-line cybersecurity and networking tool built with Go.

The goal of this project is to explore fundamental cybersecurity and networking concepts by building practical security utilities from scratch.

## Current Features

- Generate SHA-256 hashes for files
- Create trusted SHA-256 baselines for files
- Verify file integrity against a saved baseline
- Detect unexpected file modifications
- Perform DNS lookups
- Display IPv4 and IPv6 addresses associated with a domain
- Check whether a specific TCP port is open or unreachable

## Usage

### Generate a SHA-256 Hash

Generate the SHA-256 hash of a file:

```bash
go run . hash <filename>
```

Example:

```bash
go run . hash test.txt
```

Example output:

```text
File: test.txt
SHA-256: <generated SHA-256 hash>
```

### Create a File Integrity Baseline

Save the current SHA-256 hash of a file as a trusted baseline:

```bash
go run . baseline <filename>
```

Example:

```bash
go run . baseline test.txt
```

Example output:

```text
Baseline created for: test.txt
```

Sentinel stores the trusted hash in:

```text
sentinel-baseline.txt
```

### Check File Integrity

Compare a file against its saved baseline:

```bash
go run . check <filename>
```

Example:

```bash
go run . check test.txt
```

If the file has not changed:

```text
File integrity verified: test.txt
```

If the file has been modified:

```text
WARNING: File has been modified: test.txt
```

### DNS Lookup

Look up the IP addresses associated with a domain:

```bash
go run . dns <domain>
```

Example:

```bash
go run . dns google.com
```

Example output:

```text
Domain: google.com
IP addresses:
- 142.250.xxx.xxx
- 2607:f8b0:...
```

Depending on the domain, Sentinel may return both IPv4 and IPv6 addresses.

### TCP Port Check

Check whether a specific TCP port on a host accepts a connection:

```bash
go run . port <host> <port>
```

Example:

```bash
go run . port example.com 443
```

Example output for an open port:

```text
Checking: example.com:443
Status: OPEN
```

If Sentinel cannot establish the TCP connection within the timeout:

```text
Checking: example.com:81
Status: CLOSED or unreachable
```

The port command checks one user-specified TCP port at a time.

## How It Works

### SHA-256 Hashing

Sentinel reads the contents of a file and generates a SHA-256 hash.

A SHA-256 hash acts like a digital fingerprint. Even a small change to the contents of a file produces a different hash.

### File Integrity Checking

When a baseline is created, Sentinel calculates the file's SHA-256 hash and saves it as a trusted value.

When the `check` command is run, Sentinel:

1. Reads the saved baseline.
2. Calculates the file's current SHA-256 hash.
3. Compares the current hash with the trusted hash.
4. Reports whether the file has changed.

This demonstrates the basic concept behind file integrity monitoring, which can be used to identify unexpected or unauthorized changes to files.

### DNS Lookup

DNS (Domain Name System) translates human-readable domain names into IP addresses that computers use to communicate across networks.

When the `dns` command is run, Sentinel performs a DNS lookup for the supplied domain and displays the IP addresses returned by the lookup.

### TCP Port Checking

Network services listen for connections on numbered ports.

For example, common ports include:

- Port 22 — SSH
- Port 80 — HTTP
- Port 443 — HTTPS

When the `port` command is run, Sentinel attempts to establish a TCP connection to the specified host and port.

If the connection succeeds, Sentinel reports the port as open. If the connection fails or times out, Sentinel reports the port as closed or unreachable.

A three-second connection timeout prevents Sentinel from waiting indefinitely for a response.

## Responsible Use

The networking features in Sentinel are intended for educational purposes, network troubleshooting, and testing systems you own or have permission to test.

## Planned Improvements

- Refactor commands into separate functions and files
- Improve command-line argument validation
- Improve error handling
- Add automated tests
- Improve generated baseline storage
- Add additional networking utilities

## Technologies and Concepts

- Go
- Git
- GitHub
- SHA-256
- File integrity monitoring
- DNS
- IPv4 and IPv6
- TCP
- Network ports
- Connection timeouts