package filepipeline
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ProcessHex(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			if value, err := strconv.ParseInt(words[i-1], 16, 64); err == nil {
				words[i-1] = strconv.FormatInt(value, 10)
			}
			// Remove the "(hex)" tag
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index after removal
		}
	}
	return strings.Join(words, " ")
}

func ProcessBin(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			if value, err := strconv.ParseInt(words[i-1], 2, 64); err == nil {
				words[i-1] = strconv.FormatInt(value, 10)
			}
			// Remove the "(bin)" tag
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index after removal
		}
	}
	return strings.Join(words, " ")
}

func ProcessText(text string) string {
	text = ProcessHex(text)
	text = ProcessBin(text)
	return text
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <inputfile> <outputfile>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Process the text
	result := ProcessText(string(content))

	// Write to output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Processing completed successfully! Output written to: %s\n", outputFile)
}
