package commands

import (
	"errors"
	"fmt"
	"github.com/millenc/golatch"
	"github.com/spf13/cobra"
	"time"
)

//Flag variables
var AccountID string
var NoOTP bool
var Silent bool

//Flag initialization
func init() {
	StatusCmd.PersistentFlags().StringVarP(&AccountID, "account", "i", "", "Account ID")
	StatusCmd.PersistentFlags().BoolVarP(&NoOTP, "nootp", "n", false, "No OTP")
	StatusCmd.PersistentFlags().BoolVarP(&Silent, "silent", "l", false, "Silent (requires SILVER, GOLD or PLATINUM subscription)")
	StatusCmd.PersistentFlags().BoolVarP(&Bare, "bare", "b", false, "Bare output (print only essential information, useful when handling the results in shell scripts for example)")
}

//Status command
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gets the current status of an account using it's account ID (--account).",
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--account)."))
		}

		if resp, err := Latch.Status(AccountID, NoOTP, Silent); err == nil {
			//Exit code
			if resp.Status() != golatch.LATCH_STATUS_ON {
				Session.ExitCode = 1
			}

			//Output
			var bareOutput string
			if Bare {
				bareOutput += resp.Status()
			} else {
				Session.AddSuccess("account is " + resp.Status())
			}

			//Two factor
			TwoFactor := resp.TwoFactor()
			if TwoFactor.Token != "" {
				if Bare {
					bareOutput += fmt.Sprintf(":%s:%d", TwoFactor.Token, TwoFactor.Generated)
				} else {
					Session.AddInfo("two factor info:\t")
					Session.AddInfo("token\t" + TwoFactor.Token)
					Session.AddInfo(fmt.Sprintf("generated\t%d (%s)", TwoFactor.Generated, time.Unix(TwoFactor.Generated/1000, 0)))
				}
			}

			if Bare {
				Session.OutputAndExit(bareOutput)
			}
		} else {
			Session.Halt(err)
		}
	},
}
