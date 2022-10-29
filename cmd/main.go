package main

import (
	"github.com/nilsyadv/baston-eventos/config"
	"github.com/nilsyadv/baston-eventos/internal/db"
	"github.com/nilsyadv/baston-eventos/server"
)

func main() {
	// viper.AutomaticEnv()
	// environment := viper.GetString("CONFIG_ENV")
	config.Init("baston_eventos")

	// initiated db instance
	db.Initdb()
	server.InitiateServer()
}
