package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input_test_1.txt")
	if err != nil {
		return
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
}
