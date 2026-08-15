# Sentinel CLI

Sentinel CLI is a command-line cybersecurity tool built with Go.

The goal of this project is to explore fundamental cybersecurity concepts by building practical security utilities from scratch.

## Current Features

- Generate a SHA-256 hash for a file
- Detect file changes by comparing hashes

## Usage

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

## How It Works

Sentinel reads the contents of a file and generates a SHA-256 hash.

A SHA-256 hash acts like a digital fingerprint for a file. If the contents of the file change, even slightly, the resulting hash will also change.

This can be used to verify file integrity and detect unexpected modifications to files.