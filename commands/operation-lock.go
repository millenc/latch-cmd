package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	OperationLockCmd.PersistentFlags().StringVarP(&AccountID, "accountid", "i", "", "Account ID")
	OperationLockCmd.PersistentFlags().StringVarP(&OperationID, "operationid", "o", "", "Operation ID")
}

//Operation lock command
var OperationLockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Locks the operation using it's account ID (--accountid) and operation ID (--operationid).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--accountid)."))
		}
		if OperationID == "" {
			Session.Halt(errors.New("You must provide an Operation ID (--operationid)."))
		}

		if err := Latch.LockOperation(AccountID, OperationID); err == nil {
			Session.AddSuccess("Operation locked!")
		} else {
			Session.Halt(err)
		}
	},
}
