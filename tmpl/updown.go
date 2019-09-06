package tmpl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
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

	log.Println("Sad to see you go :(")
	dataDir := DefaultDataDir()

	if doesExist(dataDir) {
		err := os.RemoveAll(dataDir)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Data directory %q removed", dataDir)
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
		log.Printf("Config file %q removed", cfgFile)
	}
}

// Setup creates all files & folders used by tmpl.
// Should be called directly after installing.
func Setup() {
	dataDir := DefaultDataDir()

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

		log.Printf("Data directory %q created", dataDir)
	} else {
		log.Println("Data directory already exists")
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
		log.Printf("Config file %q created", cfgFile)
	} else {
		log.Println("Config file already exists")
	}

	log.Println("You're good to go!")
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
