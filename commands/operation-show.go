package commands

import (
	"github.com/spf13/cobra"
)

//Flag initialization
func init() {
	OperationShowCmd.PersistentFlags().StringVarP(&OperationID, "operation", "o", "", "Operation ID")
}

//Show operation command
var OperationShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows information about an operation",
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := Latch.ShowOperation(OperationID); err == nil {
			id, operation := resp.FirstOperation()

			Session.AddSuccess("operation info:\t")
			Session.AddInfo("id\t" + id)
			Session.AddInfo("name\t" + operation.Name)
			Session.AddInfo("two factor\t" + operation.TwoFactor)
			Session.AddInfo("lock on request\t" + operation.LockOnRequest)
		} else {
			Session.Halt(err)
		}
	},
}
