package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func editFile(file string) {
	editor := os.Getenv("EDITOR")
	baseName := file
	fullPath := filepath.Join(templatesDir(), baseName)

	cmd := exec.Command(editor, fullPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func copyFile(from string, to string) error {
	bytes, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}
	if _, err := os.Stat(to); os.IsNotExist(err) {
		err = ioutil.WriteFile(to, bytes, os.FileMode(0644))
		if err != nil {
			return nil
		}
	}
	return fmt.Errorf("File %q already exists", to)
}

func overwriteFile(from string, to string) error {
	err := os.Remove(to)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(to, bytes, os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}

// return a string slice with the names of all files in a directory
func readDir(path string) []string {
	var fileNames []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}
	}

	return fileNames
}

// returns path to the directory containing a user's templates
func templatesDir() string {
	home := os.Getenv("HOME")
	cfg := os.Getenv("XDG_CONFIG_HOME")
	temple := os.Getenv("TMPL_DIR")

	if temple != "" {
		return temple
	}

	if home == "" {
		home = "~"
	}
	if cfg == "" {
		cfg = filepath.Join(home, ".config")
	}

	temple = filepath.Join(cfg, "tmpl", "templates")

	if _, err := os.Stat(temple); os.IsNotExist(err) {
		_ = os.MkdirAll(temple, 0777)
	}

	return temple
}

func editor() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return "vim"
	}
	return editor
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
