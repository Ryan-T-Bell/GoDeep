package forge

import (
    "fmt"
    "os/exec"
    "os"
	"godeep/server/utils"
)

func parseForgeInput(input string) *exec.Cmd {
	cmd := exec.Command("env", "GOOS=darwin", "GOARCH=amd64", "go", "build", "-o", utils.GetOutputPath("macos"), "godeep/implant")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func Generate(input string) {
	cmd := parseForgeInput(input)
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}