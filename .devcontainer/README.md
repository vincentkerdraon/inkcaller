# InkCaller Development Container

This devcontainer provides a complete development environment for the InkCaller Go project.

## Features

- **Go 1.23**: Latest Go runtime matching the project requirements
- **VS Code Extensions**: Pre-configured Go development extensions
- **Development Tools**: Git, GitHub CLI, and Go tools
- **Custom Aliases**: Convenient shortcuts for common Go commands
- **Auto-formatting**: Code formatting on save with goimports
- **Testing Support**: Integrated test runner and explorer

## Included VS Code Extensions

- `golang.go` - Official Go extension
- `ms-vscode.vscode-json` - JSON language support
- `ms-vscode.test-adapter-converter` - Test adapter converter
- `hbenl.vscode-test-explorer` - Test explorer UI

## Custom Aliases

The container includes helpful aliases for Go development:

- `gob` - `go build`
- `gor` - `go run`
- `got` - `go test`
- `gotv` - `go test -v`
- `goc` - `go clean`
- `gom` - `go mod`
- `gof` - `go fmt`
- `goi` - `go install`
- `test-all` - `go test ./...`
- `test-verbose` - `go test -v ./...`
- `build-all` - `go build ./...`
- `lint` - `golangci-lint run`

## Getting Started

1. Open the project in VS Code
2. When prompted, click "Reopen in Container"
3. Wait for the container to build and initialize
4. Start developing!

## Port Forwarding

The container automatically forwards ports 8080 and 3000 for web services.

## Project Structure

The workspace includes:
- Main module: `github.com/vincentkerdraon/inkcaller`
- Library module: `./inkcallerlib`
- Go workspace configuration with both modules

## Environment Variables

- `GO111MODULE=on` - Module mode enabled
- `GOPROXY=https://proxy.golang.org,direct` - Go module proxy
- `GOSUMDB=sum.golang.org` - Go checksum database
