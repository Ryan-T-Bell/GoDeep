package server

/*
	This is the starting point for Go-DEEP.
	This calls the RunCLI() function, a loop
	to read stdinput and call functions based on input.
*/

import (
	"godeep/cli"
)

func main() {
	cli.RunCLI()
}
