//go:build tools
// +build tools

//go:generate go build -o ../bin/cobra-cli github.com/spf13/cobra-cli
//go:generate go build -o ../bin/wire github.com/google/wire/cmd/wire

// Package tools contains go:generate commands for all project tools with versions stored in local go.mod file
// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/spf13/cobra-cli"
)
