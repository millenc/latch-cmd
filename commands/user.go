package commands

import (
	"github.com/millenc/golatch"
	"github.com/millenc/latch-cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//Flag variables
var UserID string

//Flag & commands initialization
func init() {
	//Flags
	UserCmd.PersistentFlags().StringVarP(&UserID, "user", "u", "", "User ID")
	UserCmd.PersistentFlags().StringVarP(&SecretKey, "secret", "s", "", "User secret key")
	UserCmd.PersistentFlags().StringVarP(&Proxy, "proxy", "p", "", "Proxy URL")
	UserCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display additional information about what's going on on each call")
	UserCmd.PersistentFlags().BoolVarP(&NoShadow, "no-shadow", "w", false, "Don't hide secret keys")

	//Bind flags to config
	viper.BindPFlag("user", UserCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("user_secret", UserCmd.PersistentFlags().Lookup("secret"))
	viper.BindPFlag("proxy", UserCmd.PersistentFlags().Lookup("proxy"))

	//Subcommands
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
			if l, err := util.NewLatchUser(viper.GetString("user"), viper.GetString("user_secret"), viper.GetString("proxy"), Verbose, util.ShowRequestInfoFn(Session, NoShadow), util.ShowResponseInfoFn(Session)); err != nil {
				Session.Halt(err)
			} else {
				LatchUser = l
			}
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		Session.End()
	},
}
