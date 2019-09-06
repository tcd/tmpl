package cmd

import (
	"log"
	"os"

	"github.com/AlecAivazis/survey"
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

		var name string
		prompt := &survey.Select{
			Message: "Choose a template to remove:",
			Options: templates.Names(),
		}
		survey.AskOne(prompt, &name)

		err = templates.Remove(name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Template %q removed", name)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	rmCmd.Flags().BoolP("noconfirm", "Y", false, "Bypass any and all confirmation messages.")
}
