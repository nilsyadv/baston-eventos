package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/baston-eventos/config"
	"github.com/nilsyadv/baston-eventos/log"
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

	server := &http.Server{
		Addr:         config.Conf.GetString("server.addr") + ":" + config.Conf.GetString("server.port"),
		Handler:      route,
		ReadTimeout:  time.Duration(config.Conf.GetInt64("server.read_time_out")) * time.Second,
		WriteTimeout: time.Duration(config.Conf.GetInt64("server.write_time_out")) * time.Second,
	}

	// server started
	log.Log.Println("server starting....")
	log.Log.Println("server has initiated on port number [:" + config.Conf.GetString("server.port") + "] ....")
	server.ListenAndServe()
}
