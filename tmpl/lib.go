package tmpl

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// log.Fatal if the error isn't nil.
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Open a file in a text editor.
func editFile(pathToFile string) {
	editor := os.Getenv("EDITOR")

	cmd := exec.Command(editor, pathToFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// copy the contents of one file to a new file.
// Existing files won't be overwritten.
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
	} else {
		return fmt.Errorf("File %q already exists", to)
	}
	return nil
}

// overwrite an existing file.
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

// ReadDir returns a string slice with the names of all files in a directory.
func ReadDir(path string) []string {
	var fileNames []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}
	}

	return fileNames
}

func editor() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return "vim"
	}
	return editor
}
