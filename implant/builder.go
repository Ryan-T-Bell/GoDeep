import (
    "fmt"
    "os"
    "runtime"
)

func generate() {
    // Get the name of the executable file.
    exeName := os.Args[0]

    // Check if the executable file is already present.
    if _, err := os.Stat(exeName); err == nil {
        // The executable file is already present, so exit.
        fmt.Println("The executable file is already present.")
        return
    }

    // Create the executable file.
    err := os.MkdirAll(os.Getenv("GOPATH")+"/bin", 0755)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Build the executable file.
    go build -o exeName .

    // Check if the executable file was built successfully.
    if err := os.Stat(exeName); err != nil {
        fmt.Println(err)
        return
    }

    // The executable file was built successfully, so print a message.
    fmt.Println("The executable file was built successfully.")
}