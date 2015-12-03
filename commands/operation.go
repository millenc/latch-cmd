package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	OperationCmd.AddCommand(OperationStatusCmd)
	OperationCmd.AddCommand(OperationLockCmd)
	OperationCmd.AddCommand(OperationUnlockCmd)
}

var OperationCmd = &cobra.Command{
	Use:   "operation",
	Short: "Manages Latch operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
