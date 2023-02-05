package implant

import "fmt"

func GenerateExecutable(os int, arch int, trigger int) {
	fmt.Printf("%d %d %d", os, arch, trigger)
}
