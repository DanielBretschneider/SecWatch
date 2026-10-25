package cli

import (
	"fmt"
	"os"
)

const version = "0.1.0"

// Run handles cli arguments and executes the command
func Run() {
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
		printHelp()

	// version: Provides information on the programms version
	case "version":
		fmt.Printf("secwatch v%s\n", version)

	// default: Used for unknown commands
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

// Prints typical help message
func printHelp() {
	fmt.Printf("secwatch v%s - Linux Security Event Explorer\n", version)
	fmt.Println()
	fmt.Println("Usage: secwatch <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  help     Show this help message")
	fmt.Println("  version  Show version information")
}
