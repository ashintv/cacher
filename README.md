# GGCache

A simple TCP-based cache server written in Go.

## Installation

### Prerequisites

- Go 1.25.0 or higher

### Quick Start

```bash
# Clone the repository
git clone <repository-url>
cd cache_golang

# Build and run
make build
make run
```

The server will start on port `:3000`.

## Project Structure

```
.
├── main.go              # Application entry point
├── server.go            # TCP server implementation
├── command.go           # Command parsing logic
├── go.mod               # Go module definition
├── Makefile             # Build automation
├── bin/
│   └── ggcache         # Compiled binary
└── cache/
    ├── cache.go        # Cache implementation
    └── cacher.go       # Cache interface
```
