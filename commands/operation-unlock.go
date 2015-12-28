package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	OperationUnlockCmd.PersistentFlags().StringVarP(&AccountID, "account", "i", "", "Account ID")
	OperationUnlockCmd.PersistentFlags().StringVarP(&OperationID, "operation", "o", "", "Operation ID")
}

//Operation unlock command
var OperationUnlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlocks the operation using it's account ID (--account) and operation ID (--operation).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--account)."))
		}
		if OperationID == "" {
			Session.Halt(errors.New("You must provide an Operation ID (--operation)."))
		}

		if err := Latch.UnlockOperation(AccountID, OperationID); err == nil {
			Session.AddSuccess("operation unlocked!")
		} else {
			Session.Halt(err)
		}
	},
}
