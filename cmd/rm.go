package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/tmpl/tmpl"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an existing template",
	Run: func(cmd *cobra.Command, args []string) {
		// noconfirm, _ := cmd.Flags().GetBool("noconfirm")
		// if noconfirm {
		// 	// rmFuncNoconfirm(projects)
		// } else {
		// 	// rmFunc(projects)
		// }
		tmpl.DeleteTemplate()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	rmCmd.Flags().BoolP("noconfirm", "Y", false, "Bypass any and all confirmation messages.")
}
