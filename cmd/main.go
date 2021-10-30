package main

import (
	"github.com/Nilesh-Coherent/baston-eventos/config"
	"github.com/Nilesh-Coherent/baston-eventos/internal/db"
	"github.com/Nilesh-Coherent/baston-eventos/server"
)

func main() {
	// viper.AutomaticEnv()
	// environment := viper.GetString("CONFIG_ENV")
	config.Init("baston_eventos")

	// initiated db instance
	db.Initdb()
	server.InitiateServer()
}
