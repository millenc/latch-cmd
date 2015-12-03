package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	OperationCmd.AddCommand(OperationStatusCmd)
	OperationCmd.AddCommand(OperationLockCmd)
	OperationCmd.AddCommand(OperationUnlockCmd)
	OperationCmd.AddCommand(OperationAddCmd)
	OperationCmd.AddCommand(OperationUpdateCmd)
	OperationCmd.AddCommand(OperationDeleteCmd)
	OperationCmd.AddCommand(OperationShowCmd)
}

var OperationCmd = &cobra.Command{
	Use:   "operation",
	Short: "Manages Latch operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
