package implant

import (
	"fmt"
	"godeep/utils"
)

func GenerateExecutable(os string, arch string, trigger string) {
	fmt.Printf("%s[*] Building Implant: %s %s %s", utils.ColorBlue, os, arch, trigger)
}
