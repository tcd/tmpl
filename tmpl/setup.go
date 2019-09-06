package tmpl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// Setup tmpl.
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

		log.Println("Data directory created")
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
		log.Println("Config file created")
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
