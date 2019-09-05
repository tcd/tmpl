package tmpl

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

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
