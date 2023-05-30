package cli

/*
cli is a package to kick off GO-DEEP functions.
*/

import (
	"bufio"
	"fmt"
	"godeep/server/utils"
	"godeep/server/generate"
	"os"
	"strings"
)

const HELP = `
-----------------------------------------------------
Go-DEEP Main Help
-----------------------------------------------------
[e] (exit): Exit Go-DEEP command line interface.
[g] (generate): Generate an agent (beacon/trigger/RAT).
[ls] (list): List all agents.
`

// Reads input from the command line.
func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	utils.PrintError(err)
	return strings.TrimSuffix(input, "\n")
}

// Exits CLI application.
func handleExit() {
	fmt.Printf("%s [?]%s Are you sure you want to exit? (y/n) > ", utils.ColorGreen, utils.RESET)
	input := readInput()
	if (input == "y" || input == "yes") {
		os.Exit(0)
	}
}

// Display help instructions
func handleHelp() {
	fmt.Printf("%s", HELP)
}

// Display invalid command message
func handleInvalidCommand() {
	fmt.Printf("%s%sInvalid command. Type \"?\" \"h\" or \"help\" for instructions. \n", utils.INFO, utils.RESET)
}

// List all agents
func handleList() {
	fmt.Printf("List\n")
}

// Get first word of input string
func parseCommand(input string) string {
	return strings.Split(input, " ")[0]
}

// Main wrapper method for app command line interface.
func RunCLI() {
	for {
		fmt.Printf("%s%s%s", utils.ColorBlue, "\nGo-DEEP > ", utils.RESET)
		input := readInput()

		switch parseCommand(input) {
		case "e", "exit":
			handleExit()
		case "h", "help", "?":
			handleHelp()
		case "g", "generate":
			generate.HandleGenerate(input)
		case "ls", "list":
			handleList()
		default:
			handleInvalidCommand()
		}
	}
}




