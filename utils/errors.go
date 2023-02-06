package utils

import (
	"fmt"
	"os/exec"
)

// Error handling functions for app.
func PrintError(err error) {
	if err != nil {
		fmt.Printf("\n%s Error: %s", WARN, err)
	}
}

func GoCmdHandleErrors(err error, cmd *exec.Cmd) {
	if err != nil {
		fmt.Print("\n--- env ---\n")
		PrintError(err)
		for _, envVar := range cmd.Env {
			fmt.Printf("\n%s", envVar)
		}
	}
}
