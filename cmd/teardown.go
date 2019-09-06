package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

// teardownCmd represents the teardown command
var teardownCmd = &cobra.Command{
	Use:    "teardown",
	Short:  "Remove files & folders used by tmpl",
	Hidden: true,
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
		tmpl.Teardown()
	},
}

func init() {
	rootCmd.AddCommand(teardownCmd)
}
