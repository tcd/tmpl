package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

// lsCmd represents the ls command
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
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
