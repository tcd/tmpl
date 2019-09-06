package tmpl

import (
	"log"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	homedir "github.com/mitchellh/go-homedir"
)

// Teardown tmpl.
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
		log.Println("Data directory removed")
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
		log.Println("Config file removed")
	}
}
