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

func RunCLI() {
	for {
		printMainInputLine()
		input := readInput()

		switch input {
		case "0", "e", "exit":
			handleExit()
		case "1", "i", "implants":
			handleImplants()
		case "2", "p", "paths":
			handlePaths()
		case "h", "help":
			handleMainHelp()
		default:
			printDefaultMessage()
		}
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	utils.PrintError(err)
	return strings.TrimSuffix(input, "\n")
}

func printMainInputLine() {
	fmt.Printf("\n%s Go-DEEP > %s", utils.ColorRed, utils.ColorNone)
}

func printDefaultMessage() {
	fmt.Printf("%s Invalid command. Type \"h\" or \"help\" for instructions. \n", utils.INFO)
}

func handleExit() {
	fmt.Printf("Are you sure you want to exit? (y/n)\n")
	input := readInput()
	if input == "y" || input == "yes" {
		os.Exit(0)
	}
}

func handleMainHelp() {
	help := `Go-DEEP Main Help
-------------------------
[0] e (exit): Exit Go-DEEP command line interface.
[1] i (implants): Unsorted database of generated implants.
[2] p (paths): Paths of triggerable implants.
`
	fmt.Print(help)
}
