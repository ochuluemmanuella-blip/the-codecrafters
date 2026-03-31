package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("********* CALCULATOR *********")
	fmt.Println("Type 'help' for commands or 'quit' to exit.")

	for {
		var op string
		fmt.Print("Operation (1:add, 2:sub, 3:div, 4:mul, 5:help): ")
		fmt.Scanln(&op)

		op = strings.ToLower(strings.TrimSpace(op))

		if op == "quit" || op == "q" {
			fmt.Println("Goodbye!")
			return
		}

		if op == "5" || op == "help" {
			printHelp()
			continue
		}
		var num1 float64
		for {
			fmt.Print("First number: ")
			var input1 string
			fmt.Scanln(&input1)
			input1 = strings.TrimSpace(input1)

			if input1 == "quit" || input1 == "q" {
				fmt.Println("Goodbye!")
				return
			}

			if input1 == "" {
				fmt.Println("Please enter a number.")
				continue
			}

			var err error
			num1, err = strconv.ParseFloat(input1, 64)
			if err != nil {
				fmt.Println("Invalid number! Please enter a valid number.")
				continue
			}
			break
		}
		var num2 float64
		for {
			fmt.Print("Second number: ")
			var input2 string
			fmt.Scanln(&input2)
			input2 = strings.TrimSpace(input2)

			if input2 == "quit" || input2 == "q" {
				fmt.Println("Goodbye!")
				return
			}

			if input2 == "" {
				fmt.Println("Please enter a number.")
				continue
			}

			var err error
			num2, err = strconv.ParseFloat(input2, 64)
			if err != nil {
				fmt.Println("Invalid number! Please enter a valid number.")
				continue
			}
			break
		}
		switch op {
		case "1", "add":
			fmt.Printf("Result: %.2f\n\n", num1+num2)

		case "2", "sub", "subtract":
			fmt.Printf("Result: %.2f\n\n", num1-num2)

		case "3", "div", "divide":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed!")
				continue
			}
			fmt.Printf("Result: %.2f\n\n", num1/num2)

		case "4", "mul", "multiply":
			fmt.Printf("Result: %.2f\n\n", num1*num2)

		default:
			fmt.Println("Invalid operation! Use 1-5 or type 'help'")
		}
	}
}

func printHelp() {
	fmt.Println("\n***** Available Commands *****")
	fmt.Println("1 or add       > Addition")
	fmt.Println("2 or sub       > Subtraction")
	fmt.Println("3 or div       > Division")
	fmt.Println("4 or mul       >  Multiplication")
	fmt.Println("5 or help      > Show this help")
	fmt.Println("quit or q      > Exit the program")
	fmt.Println("************************************")
}
