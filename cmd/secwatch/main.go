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
