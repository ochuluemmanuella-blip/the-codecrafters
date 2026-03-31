package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("***************Welcome to my BASE converter****************")
	fmt.Println("type 'quit' OR 'q' to exit")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "quit" || input == "q" {
			break
		}
		parts := strings.Fields(input)

		if len(parts) != 3 {
			fmt.Println("Invalid format")
			continue
		}
		comd := parts[0]
		value := parts[1]
		base := parts[2]
		if comd != "convert" || comd != "conv" {
			fmt.Println("Invalid command")
			continue
		}
		var baseN int
		switch base {
		case "bin":
			baseN = 2
		case "hex":
			baseN = 16
		case "dec":
			baseN = 10
		default:
			fmt.Println("Invalid base")
			continue
		}
		num, err := strconv.ParseInt(value, baseN, 64)
		if err != nil {
			fmt.Println("Invalid number for base")
			continue
		}
		if base == "dec" {
			bin := strconv.FormatInt(num, 2)
			hex := strings.ToUpper(strconv.FormatInt(num, 16))

			fmt.Println("Binary: ", bin)
			fmt.Println("Hex: ", hex)
		} else {
			fmt.Println("Decimal: ", num)
		}
	}
}
