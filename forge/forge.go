package forge

import (
    "fmt"
    "os/exec"
    "godeep/implant"
)

func Generate() {
	cmd := exec.Command("go", "build", "-o", "implant.exe", "-ldflags", "-X implant.HelloWorld")
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
        return
    }
}