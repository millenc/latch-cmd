package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	LockCmd.PersistentFlags().StringVarP(&AccountID, "accountid", "i", "", "Account ID")
}

//Lock command
var LockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Locks an account using it's account ID (--accountid).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--accountid)."))
		}

		if err := Latch.Lock(AccountID); err == nil {
			Session.AddSuccess("Account Locked!")
		} else {
			Session.Halt(err)
		}
	},
}
