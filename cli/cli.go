package cli

/*
cli is a package to kick off GO-DEEP functions.
This is a command line interface
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
		case "0":
			handleExit()
		case "exit":
			handleExit()
		case "h":
			handleHelp()
		case "help":
			handleHelp()
		default:
			fmt.Printf("%s[-] Invalid command.  Type \"h\" or \"help\" for instructions. \n", utils.ColorYellow)
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
	fmt.Printf("%s Go-DEEP > %s", utils.ColorRed, utils.ColorNone)
}

func handleExit() {
	fmt.Printf("Are you sure you want to exit? (y/n)\n")
	input := readInput()
	if input == "y" || input == "yes" {
		os.Exit(0)
	}
}

func handleHelp() {
	help := `
Go-DEEP
-------------------------
[0] e (exit): Exit Go-DEEP Command Line.
[1] i (implants): Unsorted database of generated implants.
[2] p (paths): 
[3] t (triggers): 

`
	fmt.Print(help)
}
