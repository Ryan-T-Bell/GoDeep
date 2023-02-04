package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const colorRed = "\033[0;31m"
const colorNone = "\033[0m"

func RunCLI() {
	for {
		fmt.Printf("%s Go-DEEP > %s", colorRed, colorNone)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = strings.TrimSuffix(input, "\n")
		if input == "exit" {
			return
		} else {
			fmt.Println(input)
		}
	}
}
