package server

import (
	"github.com/gorilla/mux"
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

	// server := &http.Server{
	// 	Addr:         config.GetString("server.addr") + ":" + config.GetString("server.port"),
	// 	Handler:      route,
	// 	ReadTimeout:  time.Duration(config.GetInt64("server.read_time_out")) * time.Second,
	// 	WriteTimeout: time.Duration(config.GetInt64("server.write_time_out")) * time.Second,
	// }

	// // server started
	// log.Log.Println("server starting....")
	// log.Log.Println("server has initiated on port number [:" + config.GetString("server.port") + "] ....")
	// server.ListenAndServe()
}
