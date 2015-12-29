package util

import (
	"github.com/spf13/viper"
)

//Tells viper where to look for config values (config files and env variables). Flags are bound by commands.
func InitConfig() {
	//Config files
	viper.SetConfigName("latch-cmd")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.latch-cmd")
	viper.AddConfigPath("/etc/latch-cmd/")
	viper.ReadInConfig()

	//Environment variables
	viper.SetEnvPrefix("latch")
	viper.AutomaticEnv()
}
