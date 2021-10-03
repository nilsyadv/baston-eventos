package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lin-sel/baston-eventos/config"
	"github.com/lin-sel/baston-eventos/log"
)

func RegisterRouter(route *mux.Router) {

	// Register Category routes
	CreateCategoryRoute(route)

	/*
	   add more routes registration
	*/
}

// InitiateServer Used to Start Server
func InitiateServer() {
	route := mux.NewRouter()
	RegisterRouter(route)

	server := http.Server{
		Addr:         config.GetString("server.server_host") + ":" + config.GetString("server.server_port"),
		Handler:      route,
		ReadTimeout:  time.Duration(config.GetInt64("server.read_timeout")),
		WriteTimeout: time.Duration(config.GetInt64("server.write_timeout")),
	}

	// server started
	log.Log.Println("server starting....")
	log.Log.Println("server has initiated on port number [:" + config.GetString("server.server_port") + "] ....")
	server.ListenAndServe()
}
