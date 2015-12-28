package commands

import (
	"github.com/millenc/latch-cmd/session"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//Command session
var Session *session.LatchCmdSession = &session.LatchCmdSession{}

//Main command. Outputs help.
var MainCmd = &cobra.Command{
	Use:   "latch-cmd",
	Short: "Latch-cmd is an unofficial command line tool to interact with the Latch API (https://latch.elevenpaths.com/)",
	Long:  `A simple command line tool to interact with the Latch API. This tool can be used to test and debug the Latch API.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

//Adds commands and executes them
func Execute() {
	AddCommands()
	MainCmd.Execute()
}

//Adds subcommands to the root command
func AddCommands() {
	MainCmd.AddCommand(AppCmd)
	MainCmd.AddCommand(UserCmd)
	MainCmd.AddCommand(AboutCmd)
}

//Init config
func init() {
	//Config
	viper.SetConfigName("latch")
	viper.AddConfigPath("/etc/latch/")
	viper.AddConfigPath("$HOME/.latch")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	//Environment variables
	viper.SetEnvPrefix("latch")
	viper.AutomaticEnv()
}
