package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Use a template",
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
		templates, err := tmpl.GetTemplates()
		if err != nil {
			log.Fatal(err)
		}

		name := tmpl.PickTemplate("Please choose a template to use:")
		t, _ := templates.GetByName(name)
		t.Use()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
