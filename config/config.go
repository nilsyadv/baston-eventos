package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var vipr *viper.Viper

func Init(env string) {
	fmt.Println("Creating New Config Instance.....")
	NewConfig(env)
}

func NewConfig(env string) {
	conf := viper.New()
	conf.SetConfigType("json")
	conf.SetConfigName(env)
	conf.AddConfigPath("../config/")
	conf.AddConfigPath("config/")
	conf.AddConfigPath(".")
	err := conf.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		fmt.Println("Config File not found.")
	} else if err != nil {
		fmt.Println("Error:", err.Error())
	}
	vipr = conf
}

// GetString Used Get Value Using Key
func GetString(key string) string {
	return vipr.GetString(key)
}

// GetString Used Get Value(int64) Using Key
func GetInt64(key string) int64 {
	return vipr.GetInt64(key)
}

// GetBool Used Get Value(boolean) Using Key
func GetBool(key string) bool {
	return vipr.GetBool(key)
}
