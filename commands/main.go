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
	Short: "Latch-cmd is an unofficial command line tool that lets you interact with the Latch API (https://latch.elevenpaths.com/).",
	Long:  `Latch-cmd is an unofficial command line tool that lets you interact with the Latch API (https://latch.elevenpaths.com/).`,
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
	viper.SetConfigName("latch-cmd")
	viper.AddConfigPath("/etc/latch-cmd/")
	viper.AddConfigPath("$HOME/.latch-cmd")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	//Environment variables
	viper.SetEnvPrefix("latch")
	viper.AutomaticEnv()
}
