package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	ApplicationDeleteCmd.PersistentFlags().StringVarP(&Application, "app", "a", "", "Application ID")
}

//Delete operation command
var ApplicationDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an existing application",
	Run: func(cmd *cobra.Command, args []string) {
		if Application == "" {
			Session.Halt(errors.New("You must provide the ID of the application that you want to delete (--app)."))
		}

		if err := LatchUser.DeleteApplication(Application); err == nil {
			Session.AddSuccess("application succesfully deleted!")
		} else {
			Session.Halt(err)
		}
	},
}
