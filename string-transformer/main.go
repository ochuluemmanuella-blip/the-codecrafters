package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("****SENTINEL STRING TRANSFORMER--- ONLINE***")
	fmt.Println("Type 'exit' to quit")
	for {
		var input string
		fmt.Print("> ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := strings.ToLower(parts[0])
		text := ""
		if len(parts) > 1 {
			text = strings.Join(parts[:1], " ")
		}
		if text == "" && command != "exit" {
			fmt.Printf("No text provided. format: %s <text>", command)
			continue
		}

		switch command {
		case "upper":
			fmt.Println(capS(text))
		case "lower":
			fmt.Println(lowS(text))
		case "cap":
			fmt.Println(capF(text))
		case "title":
			fmt.Println(articlesT(text))
		case "snake":
			fmt.Println(underScore(text))
		case "reverse":
			fmt.Println(reverseS(text))
		case "exit":
			fmt.Println("shutting down transformer....")
			return
		default:
			fmt.Printf("Unknown Command: \"%s\"\n", command)
			fmt.Println("Commands: upper, lower, cap, title, snake, reverse, exit")
		}
	}
}
func capS(s string) string {
	return strings.ToUpper(s)
}
func lowS(s string) string {
	return strings.ToLower(s)
}
func reverseS(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		rune := []rune(w)
		for g, j := 0, len(rune)-1; g < j; g, j = g+1, j-1 {
			rune[g], rune[j] = rune[j], rune[g]
		}
		words[i] = string(rune)
	}
	return strings.Join(words, " ")
}
func capF(s string) string {
	return strings.Title(strings.ToLower(s))
}

func articlesT(s string) string {
	str := strings.Fields(s)
	seen := map[string]bool{
		"a": true, "an": true, "the": true, "and": true, "but": true, "or": true, "for": true, "nor": true, "on": true, "at": true, "to": true, "by": true, "in": true, "of": true, "up": true, "as": true, "is": true, "it": true,
	}

	for i, w := range str {
		if !seen[w] {
			str[i] = strings.Title(strings.ToLower(w))
		}
		str[0] = strings.Title(str[0])
	}
	return strings.Join(str, " ")
}
func underScore(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if !unicode.IsDigit(runes[i]) && !unicode.IsLetter(runes[i]) && runes[i] == ' ' {

			run := string(runes[:i]) + string(runes[i+1:])

			runes = []rune(run)
		}
	}
	word := strings.Fields(strings.ToLower(string(runes)))
	return strings.Join(word, "_")
}
