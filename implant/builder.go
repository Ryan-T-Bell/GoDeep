package implant

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func getImplantDirectory() string {
	user, _ := user.Current()
	dir := filepath.Join(user.HomeDir, ".godeep-implants")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

// GetGoRootDir - Get the path to GOROOT
func getGoRootDir(appDir string) string {
	return filepath.Join(appDir, goDirName)
}

// GoCmd - Execute a go command
func goCmd(config implantConfig, cwd string, command []string) ([]byte, error) {
	goBinPath := filepath.Join(getGoRootDir(""), "bin", "go")
	cmd := exec.Command(goBinPath, command...)
	cmd.Dir = cwd
	cmd.Env = []string{
		fmt.Sprintf("CC=%s", config.CC),
		fmt.Sprintf("CGO_ENABLED=%s", config.CGO),
		fmt.Sprintf("GOOS=%s", config.GOOS),
		fmt.Sprintf("GOARCH=%s", config.GOARCH),
		fmt.Sprintf("GOPATH=%s", config.ProjectDir),
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Print("go cmd: '%v'", cmd)
	err := cmd.Run()
	if err != nil {
		fmt.Print("--- env ---\n")
		for _, envVar := range cmd.Env {
			fmt.Print("%s\n", envVar)
		}
	}

	return stdout.Bytes(), err
}

// GoBuild - Execute a go build command, returns stdout/error
func goBuild(config implantConfig, src string, dest string, buildmode string, tags []string, ldflags []string, gcflags, asmflags string, trimpath string) ([]byte, error) {
	target := fmt.Sprintf("%s/%s", config.GOOS, config.GOARCH)
	if _, ok := validCompilerTargets(config)[target]; !ok {
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

	return goCmd(config, src, goCommand)
}

// GoMod - Execute go module commands in src dir
func goMod(config implantConfig, src string, args []string) ([]byte, error) {
	goCommand := []string{"mod"}
	goCommand = append(goCommand, args...)
	return goCmd(config, src, goCommand)
}

// GoVersion - Execute a go version command, returns stdout/error
func goVersion(config implantConfig) ([]byte, error) {
	var goCommand = []string{"version"}
	wd, _ := os.Getwd()
	return goCmd(config, wd, goCommand)
}

// ValidCompilerTargets - Returns a map of valid compiler targets
func validCompilerTargets(config implantConfig) map[string]bool {
	validTargets := make(map[string]bool)
	for _, target := range goToolDistList(config) {
		validTargets[target] = true
	}
	return validTargets
}

// GoToolDistList - Get a list of supported GOOS/GOARCH pairs
func goToolDistList(config implantConfig) []string {
	var goCommand = []string{"tool", "dist", "list"}
	wd, _ := os.Getwd()
	data, err := goCmd(config, wd, goCommand)
	if err != nil {
		return nil
	}
	lines := strings.Split(string(data), "\n")
	return lines
}
