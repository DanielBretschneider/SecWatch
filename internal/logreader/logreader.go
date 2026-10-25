package logreader

import (
	"bufio"
	"os"
)

// ReadFiles() - reads a file and returns it's content as individual lines
func ReadFiles(path string) ([]string, error) {
	// try to open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// prepare variable for file lines
	var lines []string
	scanner := bufio.NewScanner(file)

	// scan lines from file to variable
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// error handling
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// return scanned lines
	return lines, err
}
