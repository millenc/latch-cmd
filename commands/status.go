package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag variables
var AccountID string
var NoOTP bool
var Silent bool

//Flag initialization
func init() {
	StatusCmd.PersistentFlags().StringVarP(&AccountID, "accountid", "i", "", "Account ID")
	StatusCmd.PersistentFlags().BoolVarP(&NoOTP, "nootp", "n", false, "No OTP")
	StatusCmd.PersistentFlags().BoolVarP(&Silent, "silent", "l", false, "Silent (requires SILVER, GOLD or PLATINUM subscription)")
}

//Status command
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gets the current status of an account using it's account ID (--accountid).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--accountid)."))
		}

		if resp, err := Latch.Status(AccountID, NoOTP, Silent); err == nil {
			Session.AddSuccess("Account status: " + resp.Status())
		} else {
			Session.Halt(err)
		}
	},
}
