package commands

import (
	"errors"
	"github.com/millenc/golatch"
	"github.com/spf13/cobra"
)

//Flag variables
var ParentID string
var Name string
var TwoFactor string
var LockOnRequest string

//Flag initialization
func init() {
	OperationAddCmd.PersistentFlags().StringVarP(&ParentID, "parent", "i", "", "Parent ID (must be an existing operation or application)")
	OperationAddCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "Name of the operation")
	OperationAddCmd.PersistentFlags().StringVarP(&TwoFactor, "two-factor", "t", golatch.DISABLED, "Two Factor Authentication (possible values are MANDATORY,OPT_IN and DISABLED)")
	OperationAddCmd.PersistentFlags().StringVarP(&LockOnRequest, "lock-on-request", "l", golatch.DISABLED, "Lock On Request (possible values are MANDATORY,OPT_IN and DISABLED)")
}

//Add operation command
var OperationAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new operation",
	Run: func(cmd *cobra.Command, args []string) {
		if ParentID == "" {
			Session.Halt(errors.New("You must provide the ID of the parent operation or application (--parent)."))
		}
		if Name == "" {
			Session.Halt(errors.New("You must provide the new operation's name (--name)."))
		}

		if resp, err := Latch.AddOperation(ParentID, Name, TwoFactor, LockOnRequest); err == nil {
			Session.AddSuccess("operation created with ID: " + resp.OperationId())
		} else {
			Session.Halt(err)
		}
	},
}
