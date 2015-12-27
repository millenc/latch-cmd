package commands

import (
	"errors"
	"github.com/millenc/golatch"
	"github.com/spf13/cobra"
)

var Application string

//Flag initialization
func init() {
	ApplicationUpdateCmd.PersistentFlags().StringVarP(&Application, "app", "a", "", "Application ID")
	ApplicationUpdateCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "Name of the application")
	ApplicationUpdateCmd.PersistentFlags().StringVarP(&ContactEmail, "email", "e", "", "Contact email")
	ApplicationUpdateCmd.PersistentFlags().StringVarP(&ContactPhone, "phone", "c", "", "Contact phone")
	ApplicationUpdateCmd.PersistentFlags().StringVarP(&TwoFactor, "two-factor", "t", "", "Two Factor Authentication (possible values are MANDATORY,OPT_IN and DISABLED)")
	ApplicationUpdateCmd.PersistentFlags().StringVarP(&LockOnRequest, "lock-on-request", "l", "", "Lock On Request (possible values are MANDATORY,OPT_IN and DISABLED)")
}

//Add operation command
var ApplicationUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates an existing application",
	Run: func(cmd *cobra.Command, args []string) {
		if Application == "" {
			Session.Halt(errors.New("You must provide the ID of the application that you want to update (--app)."))
		}

		applicationInfo := &golatch.LatchApplicationInfo{
			Name:          Name,
			ContactEmail:  ContactEmail,
			ContactPhone:  ContactPhone,
			TwoFactor:     TwoFactor,
			LockOnRequest: LockOnRequest,
		}

		if err := LatchUser.UpdateApplication(Application, applicationInfo); err == nil {
			Session.AddSuccess("application succesfully updated!")
		} else {
			Session.Halt(err)
		}
	},
}
