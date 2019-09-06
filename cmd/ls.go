package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List existing templates",
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
		names := templates.Names()
		if len(names) == 0 {
			log.Println("No templates")
			os.Exit(0)
		}
		for _, name := range names {
			log.Println(name)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
