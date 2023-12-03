package fileutils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadFile(inputFileName string) ([]string, error) {
	_, err := os.Stat(inputFileName)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	}

	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		// Print each line
		input = append(input, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	return input, nil
}
