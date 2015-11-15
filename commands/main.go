package commands

import (
	"errors"
	"github.com/millenc/golatch"
	"github.com/millenc/latch-cmd/session"
	"github.com/spf13/cobra"
)

//Flag variables
var AppID string
var SecretKey string

//Flag initialization
func init() {
	MainCmd.PersistentFlags().StringVarP(&AppID, "app", "a", "", "Application's ID")
	MainCmd.PersistentFlags().StringVarP(&SecretKey, "secret", "s", "", "Secret key")
}

//Latch struct
var Latch *golatch.Latch

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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//Init latch struct used by subcommands
		if cmd.Use != "latch-cmd" {
			if l, err := NewLatch(AppID, SecretKey); err != nil {
				Session.Halt(err)
			} else {
				Latch = l
			}
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		Session.Output()
	},
}

//Adds commands and executes them
func Execute() {
	AddCommands()
	MainCmd.Execute()
}

//Adds subcommands to the root command
func AddCommands() {
	MainCmd.AddCommand(PairCmd)
	MainCmd.AddCommand(UnpairCmd)
	MainCmd.AddCommand(StatusCmd)
	MainCmd.AddCommand(OperationStatusCmd)
	MainCmd.AddCommand(LockCmd)
	MainCmd.AddCommand(UnlockCmd)
}

//Initializes the latch object that will be used by all subcommands
func NewLatch(AppID string, SecretKey string) (latch *golatch.Latch, err error) {
	if AppID == "" {
		err = errors.New("You must provide an Application's ID (--appid).")
	}
	if err == nil && SecretKey == "" {
		err = errors.New("You must provide the secret key (--key).")
	}

	if err == nil {
		latch = golatch.NewLatch(AppID, SecretKey)
	}

	return latch, err
}
