package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	/*ApplicationCmd.AddCommand(ApplicationAddCmd)
	ApplicationCmd.AddCommand(ApplicationUpdateCmd)
	ApplicationCmd.AddCommand(ApplicationDeleteCmd)*/
	ApplicationCmd.AddCommand(ApplicationShowCmd)
}

//Application base command
var ApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Manages Latch applications",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
