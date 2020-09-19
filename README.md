![Go](https://github.com/onesky/onesky-sdk-cli/workflows/Go/badge.svg?branch=dev)

# onesky-sdk-cli

The OneSky command-line interface is a Go application that lets you access OneSky from the command line. You can use this tool to perform many common platform tasks either from the command line or in scripts and other automations.

#### Build from sources:

**Required: "go 1.14" (https://golang.org/dl/)**
- Get sources from gitHub:  
`git clone https://github.com/onesky/onesky-sdk-cli.git`

- Go to source dir:

`cd OneSky-cli`

- Get dependencies: 

`go get`

- Use `make`:

`make setup`

- Or run build command for specific platform:

`GOOS=<GOOS> GOARCH=<GOARCH> go build -o bin/onesky src/onesky.go`

GOOS is one of: windows, linux, darwin, android or freebsd
GOARCH is one of: amd64, 386, arm, arm64

- Change permissions:

`chmod +x bin/onesky`

- Run:

`./bin/onesky`

- Add to PATH(optional):

`export PATH=$PATH:$(pwd)"/bin/onesky"`
