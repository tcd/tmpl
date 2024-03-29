package tmpl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	homedir "github.com/mitchellh/go-homedir"
)

// Teardown removes all files & folders used by tmpl.
// Should be called before uninstalling.
func Teardown() {
	doIt := false
	prompt := &survey.Confirm{
		Message: "Are you sure you want to remove all files & folders related to tmpl?",
	}
	survey.AskOne(prompt, &doIt)
	if !doIt {
		return
	}
	blue := color.FgBlue.Render

	dataDir := DefaultDataDir()

	if doesExist(dataDir) {
		err := os.RemoveAll(dataDir)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Data directory %s removed\n", blue(dataDir))
	}

	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	cfgFile := filepath.Join(home, ".tmpl.json")
	if doesExist(cfgFile) {
		err := os.Remove(cfgFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Config file %s removed\n", blue(cfgFile))
	}
}

// Setup creates all files & folders used by tmpl.
// Should be called directly after installing.
func Setup() {
	dataDir := DefaultDataDir()
	blue := color.FgBlue.Render

	// Make data dir
	if !doesExist(dataDir) {
		err := os.MkdirAll(dataDir, os.FileMode(0777))
		if err != nil {
			log.Fatal(err)
		}

		err = os.MkdirAll(filepath.Join(dataDir, "single"), os.FileMode(0777))
		if err != nil {
			log.Fatal(err)
		}

		err = os.MkdirAll(filepath.Join(dataDir, "multi"), os.FileMode(0777))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Data directory %s created\n", blue(dataDir))
	} else {
		fmt.Println("Data directory already exists")
	}

	// Make config file
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	cfgFile := filepath.Join(home, ".tmpl.json")
	if !doesExist(cfgFile) {
		err = createConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Config file %s created\n", blue(cfgFile))
	} else {
		fmt.Println("Config file already exists")
	}

	fmt.Println("You're good to go!")
}

func createConfig(path string) error {
	cfg := CFG{TemplatesDir: DefaultDataDir()}
	bytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, bytes, os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}
