package generate

import (
	"fmt"
	"godeep/server/utils"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const EXAMPLE = `
Examples:
generate windows amd64
generate windows 386
generate linux amd64
generate linux 386
generate darwin amd64
`

// Operating System Validator
func checkGOOS(goos string) bool {
	return goos == "darwin" || goos == "linux" || goos == "windows"
}

// Architecture Validator
func checkGOARCH(goarch string) bool {
	return goarch == "386" || goarch == "amd64" || goarch == "arm"
}

// Parse forge input
func parseForgeInput(input string) (string, string) {

	words := strings.Split(input, " ")

	// Case 1: No args (generate for current build environment)
	if words == nil || len(words) == 1 {
		fmt.Printf("%s%sNo GOOS / GOARCH specified, generating implant for current build environment: \n%s%s%s, %s\n", utils.INFO, utils.RESET, utils.GOOD, utils.RESET, runtime.GOOS, runtime.GOARCH)
		return runtime.GOOS, runtime.GOARCH
	}

	// Case 2: Invalid number of arguments
	if len(words) < 3 {
		fmt.Printf("%s%sInvalid input, please specify GOOS and GOARCH%s\n", utils.WARN, utils.RESET, EXAMPLE)
		return "", ""
	}

	goos, goarch := words[1], words[2]

	// Case 3: Invalid input, return empty strings
	if !checkGOOS(goos) || !checkGOARCH(goarch) {
		fmt.Printf("%s%sInvalid input, please specify GOOS and GOARCH%s\n", utils.WARN, utils.RESET, EXAMPLE)
		return "", ""
	}

	// Case 4: Valid input, return input
	fmt.Printf("%s%sGenerating implant for specified build environment: \n%s%s%s %s\n", utils.INFO, utils.RESET, utils.GOOD, utils.RESET, goos, goarch)
	return goos, goarch
}

// Build implant for specified GOOS and GOARCH
func getBuildCommand(goos string, goarch string) *exec.Cmd {
	cmd := exec.Command("env", "GOOS="+goos, "GOARCH="+goarch, "go", "build", "-o", utils.GetImplantOutputPath(goos+"-"+goarch), "godeep/implant")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func HandleGenerate(input string) {
	goos, goarch := parseForgeInput(input)

	if goos != "" && goarch != "" {
		cmd := getBuildCommand(goos, goarch)
		err := cmd.Run()

		if err != nil {
			fmt.Println(err)
		}
	}
}
