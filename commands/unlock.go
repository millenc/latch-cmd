package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	UnlockCmd.PersistentFlags().StringVarP(&AccountID, "account", "i", "", "Account ID")
}

//Unlock command
var UnlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlocks an account using it's account ID (--account).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--account)."))
		}

		if err := Latch.Unlock(AccountID); err == nil {
			Session.AddSuccess("account unlocked!")
		} else {
			Session.Halt(err)
		}
	},
}
