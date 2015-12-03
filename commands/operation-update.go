package commands

import (
	"errors"
	"github.com/millenc/golatch"
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	OperationUpdateCmd.PersistentFlags().StringVarP(&OperationID, "operationid", "o", "", "Operation ID")
	OperationUpdateCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "Name of the operation")
	OperationUpdateCmd.PersistentFlags().StringVarP(&TwoFactor, "two-factor", "t", golatch.NOT_SET, "Two Factor Authentication (possible values are MANDATORY,OPT_IN and DISABLED)")
	OperationUpdateCmd.PersistentFlags().StringVarP(&LockOnRequest, "lock-on-request", "l", golatch.NOT_SET, "Lock On Request (possible values are MANDATORY,OPT_IN and DISABLED)")
}

//Update operation command
var OperationUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates an operation",
	Run: func(cmd *cobra.Command, args []string) {
		if OperationID == "" {
			Session.Halt(errors.New("You must provide the operation's ID (--operationid)."))
		}

		if err := Latch.UpdateOperation(OperationID, Name, TwoFactor, LockOnRequest); err == nil {
			Session.AddSuccess("Operation was updated successfully!")
		} else {
			Session.Halt(err)
		}
	},
}
