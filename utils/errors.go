package utils

import (
	"fmt"
	"os/exec"
)

// Error handling functions for app.
func PrintError(err error) {
	if err != nil {
		fmt.Printf("%s Error: %s", WARN, err)
	}
}

func GoCmdHandleErrors(err error, cmd *exec.Cmd) {
	if err != nil {
		fmt.Print("--- env ---\n")
		for _, envVar := range cmd.Env {
			fmt.Printf("%s %s\n", WARN, envVar)
		}
	}
}
