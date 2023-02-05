package utils

import (
	"fmt"
)

/*
Error handling functions for app.
*/
func PrintError(err error) {
	if err != nil {
		fmt.Printf("%s Error: %s", WARN, err)
	}
}
