package tmpl

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// doesExist returns true if a file or folder already exists.
func doesExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

// PickFile prompts a user to choose a file from a given directory.
func PickFile(dir, message string) string {
	files := ReadDir(dir)
	if len(files) == 0 {
		log.Fatal("ReadDir: no files in ", dir)
	}

	fileName := ""
	err := survey.AskOne(
		&survey.Select{
			Message: message,
			Options: ReadDir(dir),
		},
		&fileName,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		log.Fatal(err)
	}
	return fileName
}

// CreateOrOverwrite writes to a file. Create if needed, truncate if necessary.
// See: https://godoc.org/os#example-OpenFile--Append
func CreateOrOverwrite(path string, content []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return err
	}
	if _, err := f.WriteAt(content, 0); err != nil {
		f.Close() // ignore error; Write error takes precedence
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// EditFile opens a file in a text editor.
func EditFile(pathToFile string) {
	editor := viper.GetString("editor")

	cmd := exec.Command(editor, pathToFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// StringToFile writes a string to a new or existing file.
// Existing files will overwritten.
func StringToFile(content string, destination string) error {
	if _, err := os.Stat(destination); !os.IsNotExist(err) {
		err = os.Remove(destination)
		if err != nil {
			return err
		}
	}
	err := ioutil.WriteFile(destination, []byte(content), os.FileMode(0644))
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
		log.Fatal(err)
	}
	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}
	}
	return fileNames
}

// DefaultDataDir returns path to the directory containing a user's templates.
func DefaultDataDir() string {
	temple := os.Getenv("TMPL_DIR")
	if temple == "" {
		cfg := os.Getenv("XDG_CONFIG_HOME")
		if cfg == "" {
			home, err := homedir.Dir()
			if err != nil {
				log.Fatal(err)
			}
			cfg = filepath.Join(home, ".config")
		}
		temple = filepath.Join(cfg, "tmpl", "templates")
	}
	return temple
}
