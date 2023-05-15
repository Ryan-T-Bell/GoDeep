package forge

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type GoConfig struct {
	ProjectDir string

	GOOS       string
	GOARCH     string
	GOROOT     string
	GOCACHE    string
	GOMODCACHE string
}

// GetGoRootDir - Get the path to GOROOT
func GetGoRootDir(appDir string) string {
	return filepath.Join(appDir, goDirName)
}

// GetGoCache - Get the OS temp dir (used for GOCACHE)
func GetGoCache(appDir string) string {
	cachePath := filepath.Join(GetGoRootDir(appDir), "cache")
	os.MkdirAll(cachePath, 0700)
	return cachePath
}

// GetGoModCache - Get the GoMod cache dir
func GetGoModCache(appDir string) string {
	cachePath := filepath.Join(GetGoRootDir(appDir), "modcache")
	os.MkdirAll(cachePath, 0700)
	return cachePath
}

// Garble requires $HOME to be defined, if it's not set we use the os temp dir
func getHomeDir() string {
	home := os.Getenv("HOME")
	if home == "" {
		return os.TempDir()
	}
	return home
}

// GoCmd - Execute a go command
func GoCmd(config GoConfig, cwd string, command []string) ([]byte, error) {
	goBinPath := filepath.Join(config.GOROOT, "bin", "go")
	cmd := exec.Command(goBinPath, command...)
	cmd.Dir = cwd
	cmd.Env = []string{
		fmt.Sprintf("GOOS=%s", config.GOOS),
		fmt.Sprintf("GOARCH=%s", config.GOARCH),
		fmt.Sprintf("GOPATH=%s", config.ProjectDir),
		fmt.Sprintf("GOCACHE=%s", config.GOCACHE),
		fmt.Sprintf("GOMODCACHE=%s", config.GOMODCACHE),
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		for _, envVar := range cmd.Env {
			fmt.Println("%s\n", envVar)
		}
	}

	return stdout.Bytes(), err
}

func GoBuild(config GoConfig, src string, dest string, buildmode string, tags []string, ldflags []string, gcflags, asmflags string, trimpath string) ([]byte, error) {
	
	target := fmt.Sprintf("%s/%s", config.GOOS, config.GOARCH)

	if _, ok := ValidCompilerTargets(config)[target]; !ok {
		return nil, fmt.Errorf(fmt.Sprintf("Invalid compiler target: %s", target))
	}
	var goCommand = []string{"build"}

	if 0 < len(trimpath) {
		goCommand = append(goCommand, trimpath)
	}
	if 0 < len(tags) {
		goCommand = append(goCommand, "-tags")
		goCommand = append(goCommand, tags...)
	}
	if 0 < len(ldflags) {
		goCommand = append(goCommand, "-ldflags")
		goCommand = append(goCommand, ldflags...)
	}
	if 0 < len(gcflags) {
		goCommand = append(goCommand, fmt.Sprintf("-gcflags=%s", gcflags))
	}
	if 0 < len(asmflags) {
		goCommand = append(goCommand, fmt.Sprintf("-asmflags=%s", asmflags))
	}
	if 0 < len(buildmode) {
		goCommand = append(goCommand, fmt.Sprintf("-buildmode=%s", buildmode))
	}
	goCommand = append(goCommand, []string{"-o", dest, "."}...)

	return GoCmd(config, src, goCommand)
}