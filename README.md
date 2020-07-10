# onesky-sdk-cli

#### Build from sources:

**Required: "go 1.14" (https://golang.org/dl/)**
- Get sources from gitHub:  
`git clone https://github.com/onesky/onesky-sdk-cli.git`

- Go to to source dir:

`cd OneSky-cli`

- Get dependencies: 

`go get`

- Run build command:

`GOPATH=<GOPATH> GOOS=<GOOS> GOARCH=<GOARCH> go build -o bin/onesky src/onesky.go`

GOPATH see: https://golang.org/cmd/go/#hdr-GOPATH_environment_variable

GOOS is one of: windows, linux, darwin, android or freebsd

GOARCH is one of: amd64, 386, arm, arm64

- Change permissions:

`chmod +x bin/onesky`

- Run:

`./bin/onesky`

- Add to PATH(optional):

`export PATH=$PATH:$(pwd)"/bin/onesky"`