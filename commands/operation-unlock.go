package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	OperationUnlockCmd.PersistentFlags().StringVarP(&AccountID, "accountid", "i", "", "Account ID")
	OperationUnlockCmd.PersistentFlags().StringVarP(&OperationID, "operationid", "o", "", "Operation ID")
}

//Operation unlock command
var OperationUnlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlocks the operation using it's account ID (--accountid) and operation ID (--operationid).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--accountid)."))
		}
		if OperationID == "" {
			Session.Halt(errors.New("You must provide an Operation ID (--operationid)."))
		}

		if err := Latch.UnlockOperation(AccountID, OperationID); err == nil {
			Session.AddSuccess("Operation unlocked!")
		} else {
			Session.Halt(err)
		}
	},
}
