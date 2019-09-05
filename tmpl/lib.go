package tmpl

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// log.Fatal if the error isn't nil.
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
		log.Println(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}
	}

	return fileNames
}

// TemplatesDir returns path to the directory containing a user's templates
func TemplatesDir() string {
	cfg := os.Getenv("XDG_CONFIG_HOME")
	temple := os.Getenv("TMPL_DIR")

	if temple == "" {

		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		if cfg == "" {
			cfg = filepath.Join(home, ".config")
		}

		temple = filepath.Join(cfg, "tmpl", "templates")
	}

	if _, err := os.Stat(temple); os.IsNotExist(err) {
		err = os.MkdirAll(temple, 0777)
		if err != nil {
			log.Println("Error creating tamplate directory: " + err.Error())
		}
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
