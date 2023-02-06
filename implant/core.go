package implant

import (
	"fmt"
	"godeep/utils"
)

const (
	goDirName = "go"

	kb = 1024
	mb = 1024 * kb
	gb = 1024 * mb
)

type implantConfig struct {
	ProjectDir string
	GOOS       string
	GOARCH     string
	CGO        string
	CC         string
	CXX        string
	Trigger    string
}

func GenerateExecutable(os string, arch string, trigger string) {
	fmt.Printf("%s[*] Building Implant: %s, %s, %s ...\n\n", utils.ColorBlue, os, arch, trigger)

	config := implantConfig{
		ProjectDir: utils.GetImplantDirectory(),
		GOOS:       os,
		GOARCH:     arch,
		CGO:        "0",
		CC:         "",
		CXX:        "",
		Trigger:    trigger,
	}
	for i, v := range goToolDistList(config) {
		fmt.Printf("$d $s", i, v)
	}
}
