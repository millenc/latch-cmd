package commands

import (
	"github.com/millenc/golatch"
	"github.com/millenc/latch-cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//Flag variables
var AppID string
var SecretKey string
var Proxy string
var Verbose bool
var NoShadow bool
var Bare bool

//Flag & commands initialization
func init() {
	AppCmd.PersistentFlags().StringVarP(&AppID, "app", "a", "", "Application's ID")
	AppCmd.PersistentFlags().StringVarP(&SecretKey, "secret", "s", "", "Secret key")
	AppCmd.PersistentFlags().StringVarP(&Proxy, "proxy", "p", "", "Proxy URL")
	AppCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display additional information about what's going on on each call")
	AppCmd.PersistentFlags().BoolVarP(&NoShadow, "no-shadow", "w", false, "Don't hide secret keys")

	//Bind flags to config
	viper.BindPFlag("app", AppCmd.PersistentFlags().Lookup("app"))
	viper.BindPFlag("secret", AppCmd.PersistentFlags().Lookup("secret"))
	viper.BindPFlag("proxy", AppCmd.PersistentFlags().Lookup("proxy"))

	//Subcommands
	AppCmd.AddCommand(PairCmd)
	AppCmd.AddCommand(UnpairCmd)
	AppCmd.AddCommand(StatusCmd)
	AppCmd.AddCommand(LockCmd)
	AppCmd.AddCommand(UnlockCmd)
	AppCmd.AddCommand(OperationCmd)
	AppCmd.AddCommand(HistoryCmd)
}

//Latch struct
var Latch *golatch.Latch

var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "Set of commands to interact with the main application API.",
	Long:  `Set of commands to interact with the main application API.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//Init latch struct used by subcommands
		if cmd.Use != "app" && cmd.Use != "operation" {
			if l, err := util.NewLatch(viper.GetString("app"), viper.GetString("secret"), viper.GetString("proxy"), Verbose, util.ShowRequestInfoFn(Session, NoShadow), util.ShowResponseInfoFn(Session)); err != nil {
				Session.Halt(err)
			} else {
				Latch = l
			}
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		Session.End()
	},
}
