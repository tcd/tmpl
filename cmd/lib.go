package cmd

import (
	"os"
	"strings"
)

// doesExist returns true if a file or folder already exists.
func doesExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func titleString() string {
	lines := []string{
		" _                   _ ",
		"| |_ _ __ ___  _ __ | |",
		"| __| '_ ` _ \\| '_ \\| |",
		"| |_| | | | | | |_) | |",
		" \\__|_| |_| |_| .__/|_|",
		"	      |_|      ",
		"",
	}

	return strings.Join(lines, "\n")
}
