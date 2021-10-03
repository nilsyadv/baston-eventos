package main

import (
	"github.com/lin-sel/baston-eventos/config"
	"github.com/lin-sel/baston-eventos/internal/db"
	"github.com/lin-sel/baston-eventos/server"
)

func main() {
	// viper.AutomaticEnv()
	// environment := viper.GetString("CONFIG_ENV")
	config.Init("baston_eventos")

	// initiated db instance
	db.Initdb()
	server.InitiateServer()
}
