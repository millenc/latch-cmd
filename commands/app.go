package commands

import (
	"errors"
	"fmt"
	"github.com/millenc/golatch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"strings"
)

//Flag variables
var AppID string
var SecretKey string
var Proxy string
var Verbose bool
var Bare bool

//Flag & commands initialization
func init() {
	AppCmd.PersistentFlags().StringVarP(&AppID, "app", "a", "", "Application's ID")
	AppCmd.PersistentFlags().StringVarP(&SecretKey, "secret", "s", "", "Secret key")
	AppCmd.PersistentFlags().StringVarP(&Proxy, "proxy", "p", "", "Proxy URL")
	AppCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display additional information about what's going on on each call")

	//Bind flags to config
	viper.BindPFlag("app", AppCmd.PersistentFlags().Lookup("app"))
	viper.BindPFlag("secret", AppCmd.PersistentFlags().Lookup("secret"))
	viper.BindPFlag("proxy", AppCmd.PersistentFlags().Lookup("proxy"))

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
			if l, err := NewLatch(viper.GetString("app"), viper.GetString("secret"), viper.GetString("proxy")); err != nil {
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

//Initializes the latch object that will be used by all subcommands
func NewLatch(AppID string, SecretKey string, Proxy string) (latch *golatch.Latch, err error) {
	if AppID == "" {
		err = errors.New("You must provide an Application's ID (--app).")
	}
	if err == nil && SecretKey == "" {
		err = errors.New("You must provide the secret key (--secret).")
	}

	if err == nil {
		latch = golatch.NewLatch(AppID, SecretKey)

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

//Prints request information
func OnLatchRequestStart(request *golatch.LatchRequest) {
	Session.AddInfo("request:\t")
	Session.AddInfo("url\t" + request.URL.String())
	Session.AddInfo("http-method\t" + request.HttpMethod)
	Session.AddInfo("date\t" + request.GetFormattedDate())
	Session.AddInfo("params\t" + request.GetSerializedParams())
	Session.AddInfo("headers\t" + request.GetSerializedHeaders())
	Session.AddInfo("signature\t" + strings.Replace(request.GetRequestSignature(), "\n", "\n\t\t", -1))
	Session.AddInfo("signature-sha1\t" + request.GetSignedRequestSignature())
	Session.AddInfo("auth-header\t" + request.GetAuthorizationHeader() + "\n\t\t")
}

//Prints response information
func OnLatchResponseReceive(request *golatch.LatchRequest, response *http.Response, responseBody string) {
	Session.AddInfo("response:\t")
	Session.AddInfo(fmt.Sprintf("http-status\t%d", response.StatusCode))
	Session.AddInfo("body\t" + responseBody + "\n\t\t")
}
