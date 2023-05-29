package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func GetWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

func GetImplantModulePath() string {
	filePath := GetWorkingDirectory()
	return strings.Replace(filePath, "/server", "/implant", 1)
}

// GetHomeDirectory returns the home directory of the current user
func getHomeDirectory() string {
	user, _ := user.Current()
	return user.HomeDir
}

// Create implant output directory if Nil
func makeOutputDirectoryIfNil(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)

		if err != nil {
			fmt.Println(err)
		}
	}
}

// Output path is $HOME/.godeep
func GetOutputPath(name string) string {
	dir := filepath.Join(getHomeDirectory(), ".godeep")

	makeOutputDirectoryIfNil(dir)
	return filepath.Join(dir, name)
}
