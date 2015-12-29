package commands

import (
	"github.com/millenc/latch-cmd/util"
	"github.com/spf13/cobra"
)

//Subscription command
var SubscriptionCmd = &cobra.Command{
	Use:   "subscription",
	Short: "Gets information about your current subscription.",
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := LatchUser.Subscription(); err == nil {
			Session.AddInfo("subscription info:\t")
			Session.AddInfo("id\t" + resp.ID())
			Session.AddInfo("applications\t" + util.FormatUsageString(resp.Applications().InUse, resp.Applications().Limit))
			Session.AddInfo("users\t" + util.FormatUsageString(resp.Users().InUse, resp.Users().Limit))
			Session.AddInfo("operations:\t")
			for name, usage := range resp.Operations() {
				Session.AddInfo(name + "\t" + util.FormatUsageString(usage.InUse, usage.Limit))
			}
		} else {
			Session.Halt(err)
		}
	},
}
