package generate

import (
    "fmt"
    "os/exec"
    "os"
	"godeep/server/utils"
)

func parseForgeInput(input string) (string, string) {
	return "", ""
}

func getGenerateCommand(goos string, goarch string) *exec.Cmd {
	cmd := exec.Command("env", "GOOS=%s", "GOARCH=%s", "go", "build", "-o", utils.GetOutputPath("macos"), "godeep/implant", goos, goarch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func Generate(input string) {
	goos, goarch := parseForgeInput(input)
	cmd := getGenerateCommand(goos, goarch)
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}