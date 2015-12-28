package commands

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

//Flag variables
var OperationID string

//Flag initialization
func init() {
	OperationStatusCmd.PersistentFlags().StringVarP(&AccountID, "account", "i", "", "Account ID")
	OperationStatusCmd.PersistentFlags().StringVarP(&OperationID, "operation", "o", "", "Operation ID")
	OperationStatusCmd.PersistentFlags().BoolVarP(&NoOTP, "nootp", "n", false, "No OTP")
	OperationStatusCmd.PersistentFlags().BoolVarP(&Silent, "silent", "l", false, "Silent (requires SILVER, GOLD or PLATINUM subscription)")
}

//Status command
var OperationStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gets the current status of an operation using an account ID (--account) and an operation ID (--operation).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--account)."))
		}
		if OperationID == "" {
			Session.Halt(errors.New("You must provide an Operation ID (--operation)."))
		}

		if resp, err := Latch.OperationStatus(AccountID, OperationID, NoOTP, Silent); err == nil {
			Session.AddSuccess("operation is " + resp.Status())
			TwoFactor := resp.TwoFactor()
			if TwoFactor.Token != "" {
				Session.AddInfo("two factor info:\t")
				Session.AddInfo("token\t" + TwoFactor.Token)
				Session.AddInfo(fmt.Sprintf("generated\t%d (%s)", TwoFactor.Generated, time.Unix(TwoFactor.Generated/1000, 0)))
			}
		} else {
			Session.Halt(err)
		}
	},
}
