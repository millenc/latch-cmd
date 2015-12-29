package commands

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/millenc/latch-cmd/util"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	t "time"
)

//Flag variables
var From string
var To string

//Flag initialization
func init() {
	HistoryCmd.PersistentFlags().StringVarP(&AccountID, "account", "i", "", "Account ID")
	HistoryCmd.PersistentFlags().StringVarP(&From, "from", "f", "", "From (date)")
	HistoryCmd.PersistentFlags().StringVarP(&To, "to", "t", "", "To (date)")
}

//Status command
var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Gets history information about an account. You can filter events between the --from and --to dates.",
	Run: func(cmd *cobra.Command, args []string) {
		var FromDate, ToDate t.Time
		var err error

		if AccountID == "" {
			Session.Halt(errors.New("You must provide an Account ID (--account)."))
		}

		if FromDate, err = util.ParseCmdDate(From); err != nil {
			Session.Halt(errors.New("The 'from' date has an incorrect format (please use dd-mm-yyyy hh:ii:ss)"))
		}
		if ToDate, err = util.ParseCmdDate(To); err != nil {
			Session.Halt(errors.New("The 'to' date has an incorrect format (please use dd-mm-yyyy hh:ii:ss)"))
		}

		if resp, err := Latch.History(AccountID, FromDate, ToDate); err == nil {
			//Write the data table
			buffer := &bytes.Buffer{}
			table := tablewriter.NewWriter(buffer)
			table.SetHeader([]string{"Time", "Action", "What", "Was", "Value", "Name", "User Agent", "IP"})
			for _, entry := range resp.History() {
				time := t.Unix(0, entry.Time*1000000)
				table.Append([]string{time.Format("02-01-2006 15:04:05"), entry.Action, entry.What, entry.Was, entry.Value, entry.Name, entry.UserAgent, entry.IP})
			}
			table.Render()

			//Generate output
			var output string
			output += "Last seen: " + t.Unix(0, resp.LastSeen()*1000000).Format("02-01-2006 15:04:05") + ", "
			output += "Client version: [" + util.FormatClientVersions(resp.ClientVersion()) + "], "
			output += "History count: " + fmt.Sprintf("%d", resp.HistoryCount()) + "\n\n"
			output += buffer.String()
			Session.AddSuccess(output)
		} else {
			Session.Halt(err)
		}
	},
}
