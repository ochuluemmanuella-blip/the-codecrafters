// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: Ochulu Agbenu Emmanuella
// Squad: THE INTERFACE

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: The Interface
// ───────────────────────────────────────────
// Input line types:
//   Lines in ALL CAPS
//   Lines in all lowercase
//   Lines with extra leading/trailing spaces
//   Lines that are only dashes or blanks
//
// Transformation rules (in order):
//   1. Convert ALL CAPS lines to Title Case
//   2. Convert all lowercase lines to uppercase
//   3. Trim all leading and trailing whitespace
//   4. Reverse the words in any line that contains the word REVERSE
//   5. Remove lines that are only dashes or blanks
//
// Output format:
//   Header: yes → "SENTINEL FIELD REPORT — PROCESSED"
//   Line numbering format: "1. "
//   Summary block: yes (in output file)
//
// Terminal summary fields:
//   > Lines read    : [number]
//   > Lines written : [number]
//   > Lines removed : [number]
//   > Rules applied : [list your 5 rules]
// ═══════════════════════════════════════════

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func capsToTitle(line string) string {
	if line == strings.ToUpper(line) && line != "" {
		words := strings.Fields(strings.ToLower(line))
		for i, w := range words {
			words[i] = strings.ToUpper(string(w[0])) + w[1:]
		}
		return strings.Join(words, " ")
	}
	return line
}

func lowerToUpper(line string) string {
	if line == strings.ToLower(line) && line != "" {
		return strings.ToUpper(line)
	}
	return line
}

func trimWhitespace(line string) string {
	return strings.TrimSpace(line)
}

func reverseIfNeeded(line string) string {
	if !strings.Contains(line, "REVERSE") {

		words := strings.Fields(line)
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}
		return strings.Join(words, " ")
	}
	return line
}

func isOnlyDashesOrBlank(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return true
	}
	for _, r := range trimmed {
		if r != '-' {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("Input and output cannot be the same file.")
		os.Exit(1)
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("File not found: %s\n", inputFile)
		os.Exit(1)
	}
	defer file.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Cannot write to output: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	scanner := bufio.NewScanner(file)
	var processedLines []string
	linesRead := 0
	linesRemoved := 0

	for scanner.Scan() {
		originalLine := scanner.Text()
		linesRead++

		line := originalLine

		line = trimWhitespace(line)
		line = reverseIfNeeded(line)
		line = capsToTitle(line)
		line = lowerToUpper(line)

		if isOnlyDashesOrBlank(line) {
			linesRemoved++
			continue
		}

		processedLines = append(processedLines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	writer.WriteString("SENTINEL INTERFACE REPORT\n\n")

	for i, line := range processedLines {
		writer.WriteString(fmt.Sprintf("%d. %s\n", i+1, line))
	}

	writer.WriteString("\n===============================================\n")
	writer.WriteString(fmt.Sprintf("Lines read    : %d\n", linesRead))
	writer.WriteString(fmt.Sprintf("Lines written : %d\n", len(processedLines)))
	writer.WriteString(fmt.Sprintf("Lines removed : %d\n", linesRemoved))
	writer.WriteString("Rules applied : 1.capsToTitle, 2.lowerToUpper, 3.trimWhitespace, 4.reverseIfNeeded, 5.removeDashesOrBlanks\n")
	writer.WriteString("===============================================\n")

	fmt.Println("> Lines read    :", linesRead)
	fmt.Println("> Lines written :", len(processedLines))
	fmt.Println("> Lines removed :", linesRemoved)
	fmt.Println("> Rules applied : Convert ALL CAPS to Title Case, Convert lowercase to uppercase, Trim whitespace, Reverse words containing REVERSE, Remove dashes/blanks")

	if linesRead == 0 {
		fmt.Println("Input file is empty. Nothing to process.")
	}

	fmt.Printf("\nProcessing completed successfully!\nOutput written to: %s\n", outputFile)
}
