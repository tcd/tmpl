package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// IfBoolFlag executes a function if a given bool flag is true.
// `os.Exit(0)` is called if successFunc sucessfully runs.
func IfBoolFlag(cmd *cobra.Command, flagName string, successFunc func()) {
	boolFlag, err := cmd.Flags().GetBool(flagName)
	if err != nil {
		log.Fatal(err)
	}
	if boolFlag {
		successFunc()
		os.Exit(0)
	}
}

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

// =============================================================================
// String Slice Functions
// =============================================================================

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// remove returns a new slice with all given strings removed.
func remove(slice []string, toRemove ...string) []string {
	newSlice := make([]string, len(slice)-len(toRemove))
	i := 0
	for _, s := range slice {
		match := false
		for _, str := range toRemove {
			if str == s {
				match = true
				break
			}
		}
		if !match {
			newSlice[i] = s
			i++
		}
	}
	return newSlice
}
