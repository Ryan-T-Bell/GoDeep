package cli

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRunCLI(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"0", ""},
		{"1", ""},
		{"2", ""},
		{"h", ""},
		{"", ""},
	}

	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.input)
		result, _ := ReadString(reader)
		if result != testCase.expected {
			t.Errorf("Expected %q, but got %q", testCase.expected, result)
		}
	}
}

func TestReadInput(t *testing.T) string {
	testCases := []struct {
		input    string
		expected string
	}{
		{"0", ""},
		{"1", ""},
		{"2", ""},
		{"h", ""},
		{"", ""},
	}

	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.input)
		result, _ := readInput(reader)
		if result != testCase.expected {
			t.Errorf("Expected %q, but got %q", testCase.expected, result)
		}
	}
}

func TestHandleExit(t *testing.T) {
	fmt.Printf("Are you sure you want to exit? (y/n)\n")
	input := readInput()
	if input == "y" || input == "yes" {
		os.Exit(0)
	}
}
