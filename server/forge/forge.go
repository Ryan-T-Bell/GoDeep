package forge

import (
    "fmt"
    "os/exec"
    "os"
)

func Generate() {

    cmd := exec.Command("go", "build", "-o", "test.exe", "../implant")
    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}