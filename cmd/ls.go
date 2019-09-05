package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List existing templates",
	PreRun: func(cmd *cobra.Command, args []string) {
		debug, err := cmd.Flags().GetBool("debug")
		logFatal(err)
		if debug {
			log.SetFlags(log.Lshortfile)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		tmpl.ListTemplates()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
