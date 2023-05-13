package implant

import (
    "fmt"
    "os/exec"
)

func Generate() {
    cmd := exec.Command("go", "build", "-o", "implant.exe", "-ldflags", "-X main.entryPoint=helloWorld")
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func helloWord() {
    fmt.Println("Hello World!")
}