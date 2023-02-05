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
	if os != "back" {
		arch := selectArchitecture()
		if arch != "back" {
			trigger := selectTriggerMethod()
			if trigger != 0 {
				implant.GenerateExecutable(os, arch, trigger)
			}
		}
	}
}

// OS and Architectures from https://github.com/golang/go/blob/master/src/go/build/syslist.go
func selectOperatingSystem() string {
	help := `Select Operating System:
-----------------------------------------------------
[0] b (back): Return to main thread.
[1] w (windows): Generate Windows implant.
[2] l (linux): Generate Linux implant.
[3] d (darwin): Generate Darwin (Mac-OS) implant.
[4] a (android): Generate Android implant.
[5] i (ios): Generate iOS implant.
[6] aix: Generate AIX implant.
[7] dragonfly: Generate Dragonfly implant.
[8] freebsd: Generate FreeBSD implant.
[9] hurd: Generate Hurd implant.
[10] illumos: Generate Illumos implant.
[11] js: Generate JavaScript implant.
[12] nacl: Generate Native Client implant.
[13] netbsd: Generate NetBSD implant.
[14] openbsd: Generate OpenBSD implant.
[15] plan9: Generate Plan 9 implant.
[16] solaris: Generate Solaris implant.
[17] zos: Generate z/OS implant.
`
	for {
		printImplantCLI()
		fmt.Printf("%s", help)

		input := readInput()

		switch input {
		case "0", "b", "back":
			return "back"
		case "1", "w", "windows":
			return "windows"
		case "2", "l", "linux":
			return "linux"
		case "3", "d", "darwin":
			return "darwin"
		default:
			fmt.Printf("%s Invalid input.  Select operating system", utils.INFO)
		}
	}
}

func selectArchitecture() string {
	help := `Select Architecture System:
-----------------------------------------------------
[0] b (back): Return to main thread.
[1] 386
[2] amd64
[3] amd64p32
[4] arm
[5] armbe
[6] arm64
[7] arm64be
[8] loong64
[9] mips
[10] mipsle
[11] mips64
[12] mips64le
[13] mips64p32
[14] mips64p32le
[15] ppc
[16] ppc64
[17] ppc64le
[18] riscv
[19] riscv64
[20] s390
[21] s390x
[22] sparc
[23] sparc64
[24] wasm"
`
	for {
		printImplantCLI()
		fmt.Printf("%s", help)

		input := readInput()

		switch input {
		case "0", "b", "back":
			return "back"
		case "1", "386":
			return "386"
		case "2", "amd64":
			return "amd64"
		case "3", "amd64p32":
			return "amd64p32"
		case "4", "arm":
			return "arm"
		case "5", "armbe":
			return "armbe"
		case "6", "arm64":
			return "arm64"
		case "7", "arm64be":
			return "arm64be"
		case "8", "loong64":
			return "loong64"
		case "9", "mips":
			return "mips"
		case "10", "mipsle":
			return "mipsle"
		case "11", "mips64":
			return "mips64"
		case "12", "mips64le":
			return "mips64le"
		case "13", "mips64p32":
			return "mips64p32"
		case "14", "mips64p32le":
			return "mips64p32le"
		case "15", "ppc":
			return "ppc"
		case "16", "ppc64":
			return "ppc64"
		case "17", "ppc64le":
			return "ppc64le"
		case "18", "riscv":
			return "riscv"
		case "19", "riscv64":
			return "riscv64"
		case "20", "s390":
			return "s390"
		case "21", "s390x":
			return "s390x"
		case "22", "sparc":
			return "sparc"
		case "23", "sparc64":
			return "sparc64"
		case "24", "wasm":
			return "wasm"
		default:
			fmt.Printf("%s Invalid input.  Select architecture.", utils.INFO)
		}
	}
}

func selectTriggerMethod() int {
	help := `Select Trigger Method:
-----------------------------------------------------
[0] b (back): Return to main thread.
[1] s (server): Trigger via server.  Relies on running process.
[2] w (webshell): Trigger via webshell. Relies on webserver.
`
	for {
		printImplantCLI()
		fmt.Printf("%s", help)

		input := readInput()

		switch input {
		case "0", "b", "back":
			return 0
		case "1", "s", "server":
			return 1
		case "2", "w", "webshell":
			return 2
		default:
			fmt.Printf("%s Invalid input.  Select trigger method.", utils.INFO)
		}
	}
}

// [2] List Implant Functions
func listImplants() {

}
