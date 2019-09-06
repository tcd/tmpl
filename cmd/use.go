package cmd

import (
	"log"
	"os"

	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "use a template",
	Run: func(cmd *cobra.Command, args []string) {
		templates, err := tmpl.GetTemplates()
		if err != nil {
			log.Fatal(err)
		}

		var name string
		prompt := &survey.Select{
			Message: "Choose a template to use:",
			Options: templates.Names(),
		}
		survey.AskOne(prompt, &name)

		t, _ := templates.GetByName(name)
		t.Use()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
