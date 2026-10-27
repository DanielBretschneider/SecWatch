// secwatch/cmd/secwatch/main.go
// Author: db
// Date: 2026-10-27
// Version: 0.1.5
// Description: Main entry point of secwatch - Linux Security Event Explorer CLI tool

package main

import (
	"os"
	"secwatch/internal/cli"
)

/*
 * Main() - Main entry point of secwatch
 */
func main() {
	// start cli
	os.Exit(cli.Run())

}
