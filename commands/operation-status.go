package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag variables
var OperationID string

//Flag initialization
func init() {
	OperationStatusCmd.PersistentFlags().StringVarP(&AccountID, "accountid", "i", "", "Account ID")
	OperationStatusCmd.PersistentFlags().StringVarP(&OperationID, "operationid", "o", "", "Operation ID")
	OperationStatusCmd.PersistentFlags().BoolVarP(&NoOTP, "nootp", "n", false, "No OTP")
	OperationStatusCmd.PersistentFlags().BoolVarP(&Silent, "silent", "l", false, "Silent (requires SILVER, GOLD or PLATINUM subscription)")
}

//Status command
var OperationStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gets the current status of an operation using an account ID (--accountid) and an operation ID (--operationid).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--accountid)."))
		}
		if OperationID == "" {
			Session.Halt(errors.New("You must provide an Operation ID (--operationid)."))
		}

		if resp, err := Latch.OperationStatus(AccountID, OperationID, NoOTP, Silent); err == nil {
			Session.AddSuccess("Operation status: " + resp.Status())
		} else {
			Session.Halt(err)
		}
	},
}
