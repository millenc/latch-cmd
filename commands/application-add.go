package commands

import (
	"errors"
	"fmt"
	"github.com/millenc/golatch"
	"github.com/millenc/latch-cmd/util"
	"github.com/spf13/cobra"
)

//Flag variables
var ContactEmail string
var ContactPhone string

//Flag initialization
func init() {
	ApplicationAddCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "Name of the application")
	ApplicationAddCmd.PersistentFlags().StringVarP(&ContactEmail, "email", "e", "", "Contact email")
	ApplicationAddCmd.PersistentFlags().StringVarP(&ContactPhone, "phone", "c", "", "Contact phone")
	ApplicationAddCmd.PersistentFlags().StringVarP(&TwoFactor, "two-factor", "t", golatch.DISABLED, "Two Factor Authentication (possible values are MANDATORY,OPT_IN and DISABLED)")
	ApplicationAddCmd.PersistentFlags().StringVarP(&LockOnRequest, "lock-on-request", "l", golatch.DISABLED, "Lock On Request (possible values are MANDATORY,OPT_IN and DISABLED)")
	ApplicationAddCmd.PersistentFlags().BoolVarP(&Bare, "bare", "b", false, "Bare output (print only essential information, useful when handling the results in shell scripts for example)")
}

//Add application command
var ApplicationAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new application",
	Run: func(cmd *cobra.Command, args []string) {
		if Name == "" {
			Session.Halt(errors.New("You must provide the name of the new application (--name)."))
		}

		applicationInfo := &golatch.LatchApplicationInfo{
			Name:          Name,
			ContactEmail:  ContactEmail,
			ContactPhone:  ContactPhone,
			TwoFactor:     TwoFactor,
			LockOnRequest: LockOnRequest,
		}

		if resp, err := LatchUser.AddApplication(applicationInfo); err == nil {
			if Bare {
				Session.OutputAndExit(fmt.Sprintf("%s:%s", resp.AppID(), resp.Secret()))
			} else {
				Session.AddSuccess("application succesfully created!:\t")
				Session.AddInfo("app id\t" + resp.AppID())
				Session.AddInfo("secret key\t" + util.FormatSecret(resp.Secret(), NoShadow))
			}
		} else {
			Session.Halt(err)
		}
	},
}
