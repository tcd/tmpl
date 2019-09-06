package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tcd/extedit"
	"github.com/tcd/tmpl/tmpl"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing template",
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
		ts, err := tmpl.GetTemplates()
		if err != nil {
			log.Fatal(err)
		}

		name := tmpl.PickTemplate("Please choose a template to edit:")
		t, _ := ts.GetByName(name)
		input := strings.NewReader(t.Content)
		diff, err := extedit.Invoke(input)
		if err != nil {
			log.Fatal(err)
		}
		t.Content = diff.Content()
		err = ts.Update(t)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Template updated")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().Bool("content", false, "Edit a template's contents")
	editCmd.Flags().Bool("data", false, "Edit a template's data")
	editCmd.Flags().Bool("name", false, "Edit a template's name")
	editCmd.Flags().Bool("file", false, "Edit a template's filename")
}
