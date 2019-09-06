package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new template",
	PreRun: func(cmd *cobra.Command, args []string) {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			log.Fatal(err)
		}
		if debug {
			log.SetFlags(log.Lshortfile)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := tmpl.SingleFromFile()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
