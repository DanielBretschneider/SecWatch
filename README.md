
# secwatch

**secwatch** is a local Linux Security Event Explorer written in Go. Its purpose is to help me improve my Go programming and Linux skills by building a practical CLI application for collecting, parsing, searching, and analyzing security-relevant events.

<img width="1220" height="566" alt="Screenshot_20261026_194727" src="https://github.com/user-attachments/assets/27b707a8-3287-4ac1-97a6-c2c322177eb6" />

## Goals

- Explore authentication events, such as successful and failed logins.
- Detect filesystem changes, such as file and directory creation.
- Investigate command execution and Linux audit events.
- Filter and search events from the command line.
- Eventually store events in SQLite and implement basic security detection rules.

This is a learning project intended for local use. Features will be implemented incrementally, starting with a simple CLI and sample log files.

## Project Structure

```text
secwatch/
├── cmd
│   └── secwatch
│       └── main.go
├── go.mod
├── internal
│   ├── cli
│   │   └── cli.go
│   ├── logreader
│   │   └── logreader.go
│   └── source
│       └── journal.go
├── LICENSE
├── README.md
└── testdata
|   └── auth.log
└── secwatch    
```

- `cmd/secwatch/` — Application entry point.
- `internal/cli/` — Command-line interface and argument handling.
- `internal/logreader/` — File reading functionality for parsing log files.
- `internal/journal/` — Systemd journal and security event source integration.
- `testdata/` — Sample logs for development and testing.
- `go.mod` — Go module definition and dependency management.
- `secwatch` — Built ELF 64-Bit LSB executable, x86-64 (SYSV), statically linked.

Additional packages for event models, log parsing, event sources, storage, and detection rules will be introduced as the project grows.

## Development Principles

- Written in Go, using the standard library wherever practical.
- Developed and tested locally on Linux.
- Built incrementally, with a focus on understanding each component.

## Status

Early development — starting with the CLI skeleton.
