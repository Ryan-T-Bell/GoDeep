package utils

import (
	"fmt"
)

func PrintError(message string, err error) {
	if err != nil {
		fmt.Printf("\n%s%s%s", ERROR, message, err)
	}
}
