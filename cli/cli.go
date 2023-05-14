package cli

/*
cli is a package to kick off GO-DEEP functions.
Think of this as the switching board for all actions.
*/

import (
	"bufio"
	"fmt"
	"godeep/utils"
	"godeep/forge"
	"os"
	"strings"
)

// Reads input from the command line.
func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	utils.PrintError(err)
	return strings.TrimSuffix(input, "\n")
}

// Exits CLI application.
func handleExit() {
	fmt.Printf("%s [?]%s Are you sure you want to exit? (y/n) > ", utils.ColorGreen, utils.ColorNone)
	input := readInput()
	if input == "y" || input == "yes" {
		os.Exit(0)
	}
}

// Main wrapper method for app command line interface.
func RunCLI() {
	for {
		fmt.Printf("%s%s%s", utils.ColorBlue, msgMain, utils.ColorNone)
		input := readInput()

		switch input {
		case "e", "exit":
			handleExit()
		case "h", "help", "?":
			fmt.Printf("%s", msgMainHelp)
		case "g", "generate":
			forge.Generate()
		default:
			fmt.Printf("%s %s", utils.INFO, msgInvalidCommand)
		}
	}
}




