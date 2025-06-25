package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("error opening file.")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// keyword to filter line
	keyword := "important"
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLines := strings.ReplaceAll(line, keyword, "necessary")
			fmt.Printf("line %d: %s\n", lineNumber, line)
			fmt.Printf("updated line %d: %s\n", lineNumber, updatedLines)
			lineNumber++
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("error scanning file.")
		return
	}
}

// use bufio
