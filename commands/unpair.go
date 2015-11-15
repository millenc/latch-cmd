package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	UnpairCmd.PersistentFlags().StringVarP(&AccountID, "accountid", "i", "", "Account ID")
}

//Unpair command
var UnpairCmd = &cobra.Command{
	Use:   "unpair",
	Short: "Unpairs an account using it's account ID (--accountid).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--accountid)."))
		}

		if err := Latch.Unpair(AccountID); err == nil {
			Session.AddSuccess("Unpair done!")
		} else {
			Session.Halt(err)
		}
	},
}
