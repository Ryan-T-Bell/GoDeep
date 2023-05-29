package forge

import (
    "fmt"
    "os/exec"
    "os"
)
func parseForgeInput() {
	fmt.Println("Hello World!")
}

func Generate() {
	commands := [3]*exec.Cmd{
		exec.Command("env", "GOOS=windows", "GOARCH=amd64", "go", "build", "-o", "/tmp/windows.exe", "godeep/implant"),
		exec.Command("env", "GOOS=linux", "GOARCH=amd64", "go", "build", "-o", "/tmp/linux", "godeep/implant"),
		exec.Command("env", "GOOS=darwin", "GOARCH=amd64", "go", "build", "-o", "/tmp/macos", "godeep/implant")}
	
	for _, cmd := range commands {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}