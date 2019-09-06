package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an existing template",
	Run: func(cmd *cobra.Command, args []string) {
		templates, err := tmpl.GetTemplates()
		if err != nil {
			log.Fatal(err)
		}

		name := tmpl.PickTemplate("Please choose a template to remove:")
		err = templates.Remove(name)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Template %q removed", name)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	rmCmd.Flags().BoolP("noconfirm", "Y", false, "Bypass any and all confirmation messages.")
}
