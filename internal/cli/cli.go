// secwatch/internal/cli/cli.go
// Author: db
// Date: 2026-10-27
// Version: 0.1.5
// Description: Command-line interface and argument handling for secwatch

package cli

import (
	"fmt"
	"os"
	"secwatch/internal/logreader"
	"secwatch/internal/source"
	"strconv"
)

const version = "0.1.5"

// Run handles cli arguments and executes the command
func Run() int {
	// check args length
	if len(os.Args) < 2 {
		fmt.Println("Usage: secwatch <command>")
		return 2
	}

	// given command by args
	command := os.Args[1]

	// first simple cli logic
	switch command {

	// help: Lists supported commands including short explanation
	case "help", "--help", "-h":
		printHelp()
		return 0

	// version: Provides information on the programms version
	case "version", "--version", "-v":
		fmt.Printf("secwatch v%s\n", version)
		return 0

	// read: read in file to analyze
	case "read", "--read", "-r":
		if len(os.Args) < 3 {
			fmt.Println("Usage: secwatch read <file>")
		}

		lines, err := logreader.ReadFiles(os.Args[2])
		if err != nil {
			fmt.Printf("[-] Error reading file: %v\n", err)
			return 1
		}

		for _, line := range lines {
			fmt.Println(line)
		}

		return 0

	// journal: read in journald logs for inspection
	case "journal", "--journal", "-j":
		// default value of lines that get fetched from journald
		lines := 20

		// if user wants to fetch a specific number of lines
		if len(os.Args) >= 3 {
			// number of lines as string
			strLines := os.Args[2]

			// check if given characters are int
			int_lines, err := strconv.Atoi(strLines)
			if err != nil {
				fmt.Println("[-] Error: Number of lines must be integer")
				return 2
			}

			// Make sure number is a positive integer
			if int_lines <= 0 {
				fmt.Println("[-] Error: number of lines must be greater than zero")
				return 2
			}

			// assign value since conversion worked
			lines = int_lines

			// info
			fmt.Printf("[*] User set 'lines' to: %d\n", lines)
		}

		// get output
		output, err := source.ReadJournal(strconv.Itoa(lines))
		if err != nil {
			fmt.Printf("[-] Error reading journal: %v\n", err)
			return 1
		}

		// print it - (TEST)
		fmt.Print(string(output))
		return 0

	// failed: Used to quickly filter out messages that indicate a failed login attempt
	case "failed", "--failed", "-f":
		// get cmd output
		output, err := source.ReadFailedLogins()
		if err != nil {
			fmt.Printf("[-] Error reading failed login events: %w\n", err)
		}

		// check if no entries returned
		if len(output) == 0 {
			fmt.Println("[-] No matching authentication failures found.")
			return 0
		}

		// print command output
		fmt.Print(string(output))
		return 0

	// default: Used for unknown commands
	default:
		fmt.Printf("Unknown command: %s\n", command)
		return 1
	}
}

// Prints typical help message
func printHelp() {
	fmt.Printf("secwatch v%s - Linux Security Event Explorer\n", version)
	fmt.Println()
	fmt.Println("Usage: secwatch <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  help             Show this help message")
	fmt.Println("  version          Show version information")
	fmt.Println("  read             Read a log file")
	fmt.Println("  journal [lines]  Show recent system journal entries (default: 20)")
	fmt.Println("  failed           Show authentication failure events")
}
