package main

import (
	"fmt"
	"os"
)

/*
 * Main() - Main entry point of secwatch
 */
func main() {
	// check args length
	if len(os.Args) < 2 {
		fmt.Println("Usage: secwatch <command>")
		return
	}

	// given command by args
	command := os.Args[1]

	// first simple cli logic
	switch command {

	// help: Lists supported commands including short explanation
	case "help":
		fmt.Println("secwatch - Linux Security Event Explorer")
		fmt.Println()
		fmt.Println("Usage: secwatch <command>")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("  help     Show this help message")
		fmt.Println("  version  Show version information")

	// version: Provides information on the programms version
	case "version":
		fmt.Println("secwatch v0.1.0")

	// default: Used for unknown commands
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
