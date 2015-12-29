package commands

import (
	"fmt"
	"github.com/millenc/latch-cmd/util"
	"github.com/spf13/cobra"
)

//Show applications command
var ApplicationShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows information about your applications",
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := LatchUser.ShowApplications(); err == nil {
			var i int = 1
			for applicationId, applicationInfo := range resp.Applications() {
				Session.AddInfo(fmt.Sprintf("application #%d:", i))
				Session.AddInfo("name\t" + applicationInfo.Name)
				Session.AddInfo("id\t" + applicationId)
				Session.AddInfo("secret key\t" + util.FormatSecret(applicationInfo.Secret, NoShadow))
				Session.AddInfo("two factor\t" + applicationInfo.TwoFactor)
				Session.AddInfo("lock on request\t" + applicationInfo.LockOnRequest)
				Session.AddInfo("phone\t" + applicationInfo.ContactPhone)
				Session.AddInfo("email\t" + applicationInfo.ContactEmail)
				Session.AddInfo("image\t" + applicationInfo.ImageURL + "\n")

				i++
			}
		} else {
			Session.Halt(err)
		}
	},
}
