package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var vipr *viper.Viper

func init() {
	fmt.Println("Creating New Config Instance.....")
	NewConfig()
}

func NewConfig() {
	conf := viper.New()
	conf.SetConfigName(os.Getenv("config"))
	conf.SetConfigType(os.Getenv("extension"))
	conf.AddConfigPath("/deployment/config")
	err := conf.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		fmt.Println("Config File not found.")
	} else {
		fmt.Println("Error in Config file Read:", err.Error())
	}
	vipr = conf
}

// GetString Used Get Value Using Key
func GetString(key string) string {
	return vipr.GetString(key)
}
