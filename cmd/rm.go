package cmd

import (
	"log"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an existing template",
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

		name := tmpl.PickTemplate("Please choose a template to remove:")
		err = templates.Remove(name)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Template %s removed\n", color.FgBlue.Render(name))
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	rmCmd.Flags().BoolP("noconfirm", "Y", false, "Bypass any and all confirmation messages.")
}
