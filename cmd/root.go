package cmd

import (
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tcd/tmpl/tmpl"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "tmpl",
	Short: "Create templates for frequently used files and project layouts",
	PreRun: func(cmd *cobra.Command, args []string) {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			log.Fatal(err)
		}
		if debug {
			log.SetFlags(log.Lshortfile)
		}

		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			log.Fatal(err)
		}
		if version {
			cmd.Help() // TODO: Add version output.
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tmpl.json)")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Ouput debug information")

	rootCmd.Flags().BoolP("version", "v", false, "Print version information")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		cfgFile = filepath.Join(home, ".tmpl.json")

		viper.AddConfigPath(home)
		viper.SetConfigName(".tmpl")
		viper.SetConfigType("json")
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		tmpl.Setup()
		os.Exit(0)
	}
}
