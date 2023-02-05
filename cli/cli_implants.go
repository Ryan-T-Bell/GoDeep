package cli

import (
	"fmt"
	"godeep/utils"
)

func handleImplants() {
	for {
		fmt.Printf("\n%s Go-DEEP [Implants] > %s", utils.ColorBlue, utils.ColorNone)
		input := readInput()

		switch input {
		case "0", "b", "back":
			return
		case "1", "g", "generate":
		case "2", "ls", "list":
		case "h", "help":
			handleImplantHelp()
		default:
			printDefaultMessage()
		}
	}
}

func handleImplantHelp() {
	help := `Go-DEEP Implant Help
-------------------------
[0] b (back): Return to main thread.
[1] g (generate): Enter queue to generate implants.
[2] ls (list): List all implants that are generated.
`
	fmt.Print(help)
}
