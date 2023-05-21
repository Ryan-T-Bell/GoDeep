package forge

import (
    "fmt"
    "os/exec"
    "os"
)

func Generate() {

    cmd := exec.Command("go", "build", "-o", "test.exe", "implant/core.go")
    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

    // Create a new build context.
    // ctx := build.Default

    // Set the target platform to Windows.
    // ctx.GOOS = "windows"

    // Set the build mode to release.
    // ctx.BuildMode = "release"

    // Build the package.
    // err := ctx.Build(build.BuildContext{
    //     Package: packagePath,
    //     Outfile: "app.exe",
    // })