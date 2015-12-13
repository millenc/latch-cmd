package commands

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/millenc/golatch"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"strings"
	t "time"
)

//Flag variables
var From string
var To string

//Flag initialization
func init() {
	HistoryCmd.PersistentFlags().StringVarP(&AccountID, "accountid", "i", "", "Account ID")
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
			Session.Halt(errors.New("You must provide an Account ID (--accountid)."))
		}

		if FromDate, err = ParseCmdDate(From); err != nil {
			Session.Halt(errors.New("The 'from' date has an incorrect format (please use dd-mm-yyyy hh:ii:ss)"))
		}
		if ToDate, err = ParseCmdDate(To); err != nil {
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
			output += "Client version: [" + FormatClientVersions(resp.ClientVersion()) + "], "
			output += "History count: " + fmt.Sprintf("%d", resp.HistoryCount()) + "\n\n"
			output += buffer.String()
			Session.AddSuccess(output)
		} else {
			Session.Halt(err)
		}
	},
}

//Parses a date received from the command line
func ParseCmdDate(date string) (parsed t.Time, err error) {
	if date == "" {
		return t.Time{}, nil
	}

	parsed, err = t.Parse("02-01-2006 15:04:05", date)

	return parsed, err
}

//Formats the client versions received from the API
func FormatClientVersions(clientVersions []golatch.LatchClientVersion) (formatted string) {
	versions := []string{}
	for _, version := range clientVersions {
		versions = append(versions, version.Platform+" - "+version.App)
	}

	return strings.Join(versions, ",")
}
