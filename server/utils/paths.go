package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

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
func GetImplantOutputPath(name string) string {
	dir := filepath.Join(getHomeDirectory(), ".godeep", "implants")

	makeOutputDirectoryIfNil(dir)

	if strings.Contains(name, "windows") {
		return filepath.Join(dir, name+".exe")
	}

	return filepath.Join(dir, name)
}
