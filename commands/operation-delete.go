package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	OperationDeleteCmd.PersistentFlags().StringVarP(&OperationID, "operationid", "o", "", "Operation ID")
}

//Delete operation command
var OperationDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an operation",
	Run: func(cmd *cobra.Command, args []string) {
		if OperationID == "" {
			Session.Halt(errors.New("You must provide the operation's ID (--operationid)."))
		}

		if err := Latch.DeleteOperation(OperationID); err == nil {
			Session.AddSuccess("Operation was deleted!")
		} else {
			Session.Halt(err)
		}
	},
}
