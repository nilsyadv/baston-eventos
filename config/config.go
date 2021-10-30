package config

import (
	"log"

	"github.com/Nilesh-Coherent/common-service-evnt/pkg/config"
)

var Conf config.IConfig

// initiating config
func Init(env string) {
	conf, err := config.InitConfig(env)
	if err != nil {
		log.Fatalf("Failed to Initiate Config: %v", err.Error())
	}
	Conf = conf
}
