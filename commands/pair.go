package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag variables
var Token string

//Flag initialization
func init() {
	PairCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "Token provided by the user to perform the pairing")
}

//Pair command
var PairCmd = &cobra.Command{
	Use:   "pair",
	Short: "Pairs an account with the provided pairing token (--token).",
	Run: func(cmd *cobra.Command, args []string) {
		if Token == "" {
			Session.Halt(errors.New("You must provide the pairing token (--token)."))
		}

		if resp, err := Latch.Pair(Token); err == nil {
			Session.AddSuccess("Pairing done! Account ID is " + resp.AccountId())
		} else {
			Session.Halt(err)
		}
	},
}
