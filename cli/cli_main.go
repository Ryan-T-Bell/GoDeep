package cli

/*
cli is a package to kick off GO-DEEP functions.
Think of this as the switching board for all actions.
*/

import (
	"bufio"
	"fmt"
	"godeep/utils"
	"os"
	"strings"
)

// Strings
const msgMain = "\nGo-DEEP > "
const msgInvalidCommand = "Invalid command. Type \"?\" \"h\" or \"help\" for instructions. \n"
const msgMainHelp = `-----------------------------------------------------
Go-DEEP Main Help
-----------------------------------------------------
[0] e (exit): Exit Go-DEEP command line interface.
[1] i (implants): Unsorted database of generated implants.
[2] p (paths): Paths of triggerable implants.
`

// Main wrapper method for app command line interface.
func RunCLI() {
	for {
		fmt.Printf("%s%s%s", utils.ColorRed, msgMain, utils.ColorNone)
		input := readInput()

		switch input {
		case "0", "e", "exit":
			handleExit()
		case "1", "i", "implants":
			handleImplants()
		case "2", "p", "paths":
			handlePaths()
		case "h", "help", "?":
			fmt.Printf("%s", msgMainHelp)
		default:
			fmt.Printf("%s %s", utils.INFO, msgInvalidCommand)
		}
	}
}

// Reads input from the command line.
func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	utils.PrintError(err)
	return strings.TrimSuffix(input, "\n")
}

// Exits CLI application.
func handleExit() {
	fmt.Printf("%s [?]%s Are you sure you want to exit? (y/n)\n> ", utils.ColorGreen, utils.ColorNone)
	input := readInput()
	if input == "y" || input == "yes" {
		os.Exit(0)
	}
}
