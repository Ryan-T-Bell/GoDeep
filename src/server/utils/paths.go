package utils

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// Get impant directory: ~/.godeep-implants
func GetImplantDirectory() string {
	user, _ := user.Current()
	dir := filepath.Join(user.HomeDir, ".godeep-implants")
	makeImplantDirectoryIfNil(dir)
	return dir
}

// Make implant directory if it does not exist
func makeImplantDirectoryIfNil(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetGoPath() string {
	user, _ := user.Current()
	return filepath.Join(user.HomeDir, "go", "bin")
}
