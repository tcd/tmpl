package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tcd/tmpl/tmpl"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create tmpl config file & data directory.",
	PreRun: func(cmd *cobra.Command, args []string) {
		debug, err := cmd.Flags().GetBool("debug")
		logFatal(err)
		if debug {
			log.SetFlags(log.Lshortfile)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if doesExist(cfgFile) {
			cfg := viper.ConfigFileUsed()
			log.Println(cfg)
			os.Exit(0)
		} else {
			err := createConfig(cfgFile)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Config file created at: ", cfgFile)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// doesExist returns true if a file or folder already exists.
func doesExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

// Create a config file at the given path.
func createConfig(path string) error {
	cfg := tmpl.CFG{TemplatesDir: tmpl.TemplatesDir()}
	bytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, bytes, os.FileMode(0644))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
