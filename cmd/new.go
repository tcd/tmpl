package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new template",
	Run: func(cmd *cobra.Command, args []string) {
		copy, err := cmd.Flags().GetBool("copy")
		logFatal(err)
		if copy {
			tmpl.CopyToTemplate()
			os.Exit(0)
		} else {
			tmpl.NewTemplate()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().BoolP("copy", "c", false, "Create a template from an existing file")
}
