package implant

import "fmt"

func GenerateExecutable(os string, arch string, trigger int) {
	fmt.Printf("%s %s %d", os, arch, trigger)
}
