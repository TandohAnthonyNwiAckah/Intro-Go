
 # Go Learning Project
Go (or Golang) is an open-source programming language developed by Google. It is designed for simplicity, efficiency, and reliability, making it ideal for building scalable and high-performance applications. Go features a clean syntax, built-in concurrency support, and a powerful standard library, making it popular for web services, cloud infrastructure, and command-line tools.


This repository contains examples and exercises for learning the Go programming language.

## Getting Started

1. Install Go from [golang.org](https://golang.org/dl/)
2. Verify installation by running: `go version`
3. Set up your workspace (GOPATH)


## Topics Covered

- Basic syntax and data types
- Control structures
- Functions and methods
- Structs and interfaces
- Goroutines and channels
- Error handling
- Testing
- Web development
- Database operations

## GO mod

Go modules are the official dependency management solution for Go. Here's what you need to know:

1. Initialize a new module:
```bash
go mod init [module-name]
```

2. Key commands:
- `go mod tidy`: Add missing dependencies and remove unused ones
- `go mod vendor`: Create vendor directory with dependencies
- `go mod verify`: Verify dependencies have expected content
- `go mod download`: Download dependencies to local cache

3. Benefits:
- Versioning and reproducible builds
- No need for GOPATH
- Better dependency management
- Semantic versioning support


## Running Examples

To run any example:

```bash
cd path/to/IntroGo
go run main.go

OR

go run . 
```
 






 ## Build Go
To build your Go project into a binary, use:

```bash
go build
```

This will create an executable file in the current directory. You can specify the output name with:

```bash
go build -o myapp
```


### Cross-Compiling with GOOS and GOARCH

Go makes it easy to build binaries for different operating systems and architectures. Set the `GOOS` and `GOARCH` environment variables before running `go build`. For example, to build a Linux binary from macOS:

```bash
GOOS=linux GOARCH=amd64 go build -o myapp-linux
```

Common values for `GOOS` include `linux`, `windows`, and `darwin` (macOS). For `GOARCH`, common values are `amd64` and `arm64`.

