package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

// teardownCmd represents the teardown command
var teardownCmd = &cobra.Command{
	Use:    "teardown",
	Short:  "Remove files & folders used by tmpl",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		tmpl.Teardown()
	},
}

func init() {
	rootCmd.AddCommand(teardownCmd)
}
