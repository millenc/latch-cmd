package commands

import (
	"errors"
	"github.com/millenc/golatch"
	"github.com/spf13/cobra"
	"net/url"
)

//Flag variables
var UserID string

//Flag & commands initialization
func init() {
	UserCmd.PersistentFlags().StringVarP(&UserID, "user", "u", "", "User ID")
	UserCmd.PersistentFlags().StringVarP(&SecretKey, "secret", "s", "", "User secret key")
	UserCmd.PersistentFlags().StringVarP(&Proxy, "proxy", "p", "", "Proxy URL")
	UserCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display additional information about what's going on on each call")

	UserCmd.AddCommand(SubscriptionCmd)
	UserCmd.AddCommand(ApplicationCmd)
}

//Latch struct
var LatchUser *golatch.LatchUser

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Set of commands to interact with the user API (manage applications and subscription information).",
	Long:  `Set of commands to interact with the user API (manage applications and subscription information).`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//Init latch struct used by subcommands
		if cmd.Use != "user" && cmd.Use != "application" {
			if l, err := NewLatchUser(UserID, SecretKey, Proxy); err != nil {
				Session.Halt(err)
			} else {
				LatchUser = l
			}
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		Session.Output()
	},
}

//Initializes the latch object that will be used by all subcommands
func NewLatchUser(UserID string, SecretKey string, Proxy string) (latch *golatch.LatchUser, err error) {
	if UserID == "" {
		err = errors.New("You must provide the User ID (--user).")
	}
	if err == nil && SecretKey == "" {
		err = errors.New("You must provide the user secret key (--secret).")
	}

	if err == nil {
		latch = golatch.NewLatchUser(UserID, SecretKey)

		if Proxy != "" {
			if proxyUrl, err := url.Parse(Proxy); err == nil {
				latch.SetProxy(proxyUrl)
			}
		}

		if Verbose {
			latch.OnRequestStart = OnLatchRequestStart
			latch.OnResponseReceive = OnLatchResponseReceive
		}
	}

	return latch, err
}
