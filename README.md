# Sentinel CLI

Sentinel CLI is a command-line cybersecurity tool built with Go.

The goal of this project is to explore fundamental cybersecurity concepts by building practical security utilities from scratch.

## Current Features

- Generate a SHA-256 hash for a file
- Create a trusted SHA-256 baseline for a file
- Verify whether a file still matches its trusted baseline
- Detect unexpected file modifications

## Usage

### Generate a SHA-256 Hash

To generate the SHA-256 hash of a file:

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

To save the current SHA-256 hash of a file as a trusted baseline:

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

To compare a file against its saved baseline:

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

## How It Works

Sentinel uses SHA-256 hashing to create a digital fingerprint of a file.

When a baseline is created, Sentinel calculates the file's SHA-256 hash and saves it as the trusted value.

When the `check` command is run, Sentinel:

1. Reads the saved baseline.
2. Calculates the file's current SHA-256 hash.
3. Compares the current hash with the saved hash.
4. Reports whether the file has changed.

Even a small modification to the contents of a file will result in a different SHA-256 hash.

This demonstrates the basic concept behind file integrity monitoring, which can be used to detect unexpected or unauthorized changes to files.

## Planned Features

- DNS lookup
- Basic network utilities
- Improved command structure
- Automated tests
- Improved error handling

## Technologies

- Go
- Git
- GitHub
- SHA-256