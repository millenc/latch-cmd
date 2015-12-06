package commands

import (
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	OperationShowCmd.PersistentFlags().StringVarP(&OperationID, "operationid", "o", "", "Operation ID")
}

//Show operation command
var OperationShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows information about an operation",
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := Latch.ShowOperation(OperationID); err == nil {
			id, operation := resp.FirstOperation()
			output := ""
			output += "ID: " + id + "\n"
			output += "Name: " + operation.Name + "\n"
			output += "Two Factor: " + operation.TwoFactor + "\n"
			output += "Lock On Request: " + operation.LockOnRequest + "\n"

			Session.AddSuccess("Operation info:" + "\n" + output)
		} else {
			Session.Halt(err)
		}
	},
}
