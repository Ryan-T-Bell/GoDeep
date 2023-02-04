package cli

import (
	"bufio"
	"fmt"
	"godeep/utils"
	"os"
	"strings"
)

const colorRed = "\033[0;31m"
const colorNone = "\033[0m"

func RunCLI() {
	for {
		printMainInputLine()
		input := readInput()

		switch input {
		case "exit":
			handleExit()
		default:
			fmt.Println(input)
		}
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	utils.PanicIfError(err)
	return strings.TrimSuffix(input, "\n")
}

func printMainInputLine() {
	fmt.Printf("%s Go-DEEP > %s", colorRed, colorNone)
}

func handleExit() {
	fmt.Printf("Are you sure you want to exit? (y/n)")
	input := readInput()
	if input == "y" || input == "yes" {
		os.Exit(0)
	}
}
