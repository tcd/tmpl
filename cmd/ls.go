package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"
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
		long, err := cmd.Flags().GetBool("long")
		if err != nil {
			log.Fatal(err)
		}
		if long {
			listLong()
			os.Exit(0)
		} else {
			listBasic()
			os.Exit(0)
		}

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolP("long", "l", false, "Output a more detailed listing")
}

func listBasic() {
	ts, err := tmpl.GetTemplates()
	if err != nil {
		log.Fatal(err)
	}
	names := ts.Names()
	if len(names) == 0 {
		fmt.Println("No templates")
		os.Exit(0)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}

func listLong() {
	ts, err := tmpl.GetTemplates()
	if err != nil {
		log.Fatal(err)
	}
	if len(ts.T) == 0 {
		fmt.Println("No templates")
		os.Exit(0)
	}

	blue := color.FgBlue.Render
	for _, t := range ts.T {
		fmt.Printf("Name: %s, FileName: %s, Variables: %s", blue(t.Name), blue(t.FileName), blue(t.Data))
	}
}
