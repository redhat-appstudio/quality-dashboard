//go:build tools
// +build tools

package tools

// tool dependencies
import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/swaggo/swag/cmd/swag"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
)

//go:generate go build -v -o=./bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go build -v -o=./bin/gotestsum gotest.tools/gotestsum
//go:generate go build -v -o=./bin/swag github.com/swaggo/swag/cmd/swag
//go:generate go build -v -o=./bin/gofumpt mvdan.cc/gofumpt
