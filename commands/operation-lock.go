package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	OperationLockCmd.PersistentFlags().StringVarP(&AccountID, "account", "i", "", "Account ID")
	OperationLockCmd.PersistentFlags().StringVarP(&OperationID, "operation", "o", "", "Operation ID")
}

//Operation lock command
var OperationLockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Locks the operation using it's account ID (--account) and operation ID (--operation).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--account)."))
		}
		if OperationID == "" {
			Session.Halt(errors.New("You must provide an Operation ID (--operation)."))
		}

		if err := Latch.LockOperation(AccountID, OperationID); err == nil {
			Session.AddSuccess("operation locked!")
		} else {
			Session.Halt(err)
		}
	},
}
