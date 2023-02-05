package cli

import (
	"fmt"
	"godeep/implant"
	"godeep/utils"
)

// HandleImplants is the main function for the implant thread.
func handleImplants() {
	for {
		printImplantCLI()
		input := readInput()

		switch input {
		case "0", "b", "back":
			return
		case "1", "g", "generate":
			generateImplantCLI()
		case "2", "ls", "list":
			listImplants()
		case "h", "help", "?":
			handleImplantHelp()
		default:
			printDefaultMessage()
		}
	}
}

// [0] Admin Functions
func printImplantCLI() {
	fmt.Printf("\n%s Go-DEEP [Implants] > %s", utils.ColorBlue, utils.ColorNone)
}

func handleImplantHelp() {
	help := `Go-DEEP Implant Help
-----------------------------------------------------
[0] b (back): Return to main thread.
[1] g (generate): Enter queue to generate implants.
[2] ls (list): List all implants that are generated.
`
	fmt.Print(help)
}

// [1] Generate Implant Functions
func generateImplantCLI() {
	os := selectOperatingSystem()
	if os != 0 {
		arch := selectArchitecture()

		if arch != 0 {
			implant.GenerateImplant(os, arch)
		}
	}
}

func selectOperatingSystem() int {
	help := `Select Operating System:
-----------------------------------------------------
[0] b (back): Return to main thread.
[1] w (windows): Generate Windows implant.
[2] l (linux): Generate Linux implant.
[3] m (mac): Generate Mac implant.
`
	for {
		printImplantCLI()
		fmt.Printf("%s", help)

		input := readInput()

		switch input {
		case "0", "b", "back":
			return 0
		case "1", "w", "windows":
			return 1
		case "2", "l", "linux":
			return 2
		case "3", "m", "mac":
			return 3
		default:
			fmt.Printf("%s Invalid input.  Select operating systme", utils.INFO)
		}
	}
}

func selectArchitecture() int {
	help := `Select Operating System:
-----------------------------------------------------
[0] b (back): Return to main thread.
[1] amd64: AMD64 architecture.
[2] x86_64: x86_64 architecture.
[3] x86: x86 architecture.
[4] arm: ARM architecture.
[5] arm64: ARM64 architecture.
[6] mips: MIPS architecture.
[7] mips64: MIPS64 architecture.
[8] mips64le: MIPS64LE architecture.
[9] mipsle: MIPSLE architecture.
[10] ppc64: PPC64 architecture.
[11] ppc64le: PPC64LE architecture.
[12] riscv64: RISCV64 architecture.
[13] s390x: S390X architecture.
`
	for {
		printImplantCLI()
		fmt.Printf("%s", help)

		input := readInput()

		switch input {
		case "0", "b", "back":
			return 0
		case "1", "amd64":
			return 1
		case "2", "x86_64":
			return 2
		case "3", "x86":
			return 3
		case "4", "arm":
			return 4
		case "5", "arm64":
			return 5
		case "6", "mips":
			return 6
		case "7", "mips64":
			return 7
		case "8", "mips64le":
			return 8
		case "9", "mipsle":
			return 9
		case "10", "ppc64":
			return 10
		case "11", "ppc64le":
			return 11
		case "12", "riscv64":
			return 12
		case "13", "s390x":
			return 13
		default:
			fmt.Printf("%s Invalid input.  Select operating systme", utils.INFO)
		}
	}
}

// [2] List Implant Functions
func listImplants() {

}
