// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: Ochulu Agbenu Emmanuella
// Squad: THE INTERFACE

/*
// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [The Interface]
// ───────────────────────────────────────────
// Input line types:
//   [
        Lines in ALL CAPS
        Lines in all lowercase
        Lines with extra leading/trailing spaces
        Lines that are only dashes or blanks
       ]
//
// Transformation rules (in order):
//   1. [Convert ALL CAPS lines to Title Case]
//   2. [Convert all lowercase lines to uppercase]
//   3. [Trim all leading and trailing whitespace]
//   4. [Reverse the words in any line that contains the word REVERSE]
//   5. [Remove lines that are only dashes or blanks]
//
// Output format:
//   [Header: yes]
//   [Line numbering format: "1"]
//   [Summary block: yes]
//
// Terminal summary fields:
//   [Lines read    : [number]]
//   [Lines written : [number]]
//   [Lines removed : [number]]
//   [Rules applied : [list your 5 rules]]
// ═══════════════════════════════════════════
═════════════════════════════════════════
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func capsToTitle(text string) string {
	if text == strings.ToUpper(text) && text != "" {
		return strings.Title(strings.ToLower(text))
	}
	return text
}
func lowerToUpper(text string) string {
	if text == strings.ToLower(text) && text != "" {
		return strings.ToUpper(text)
	}
	return text
}
func trimSpaces(text string) string {
	return strings.Join(strings.Fields(text), " ")
}
func removeDashesAndBlanks(text string) bool {
	if text == "" {
		return true
	}
	cleaned := strings.TrimSpace(text)
	if cleaned == "" {
		return true
	}
	for _, r := range cleaned {
		if r != '-' {
			return false
		}
	}
	return true
}
func reverseS(text string) string {
	if strings.Contains(text, "REVERSE") {
		words := strings.Fields(text)
		for g, j := 0, len(words)-1; g < j; g, j = g+1, j-1 {
			words[g], words[j] = words[j], words[g]
		}
		return strings.Join(words, " ")
	}
	return text
}
func ProcessText(content string) (string, int, int, int) {
	line := strings.Split(content, "\n")
	var cleaned []string
	linesRead := 0
	linesRemoved := 0
	for _, lines := range line {
		linesRead++

		lines = capsToTitle(lines)
		lines = lowerToUpper(lines)
		lines = trimSpaces(lines)
		lines = reverseS(lines)

		if removeDashesAndBlanks(lines) {
			linesRemoved++
			continue
		}
		cleaned = append(cleaned, lines)
	}

	var output strings.Builder
	output.WriteString("==== THE INTERFACE FILE-PIPELINE OUTPUR ====\n\n")
	for i, line := range cleaned {
		output.WriteString(fmt.Sprintf("%d. %s\n", i+1, line))
	}
	output.WriteString("\n")
	output.WriteString("===============================================\n")
	output.WriteString(fmt.Sprintf("lines read : %d\n", linesRead))
	output.WriteString(fmt.Sprintf("Lines Written : %d\n", len(cleaned)))
	output.WriteString(fmt.Sprintf("Lines Removed : %d\n", linesRemoved))
	output.WriteString("Rules Applied : 1.capsToTitle, 2.lowerToUpper, 3.trimSpaces, 4.reverseS, 5.removeDashesAndBlanks")
	output.WriteString("================================================\n")

	return output.String(), linesRead, len(cleaned), linesRemoved
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <inputfile> <outputfile>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	result := ProcessText(string(content))

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Processing completed successfully! Output written to: %s\n", outputFile)
}
