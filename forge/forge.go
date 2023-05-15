package forge

import (
    "fmt"
    "os/exec"
)

func Generate() {
	cmd := exec.Command("go", "build", "-o", "test.exe", "-ldflags", "-X /godeep/implant.HelloWorld")
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
        return
    }
}