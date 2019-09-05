package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing template",
	Run: func(cmd *cobra.Command, args []string) {
		tmpl.EditTemplate()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
