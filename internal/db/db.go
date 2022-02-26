package db

import (
	"github.com/Nilesh-Coherent/baston-eventos/config"
	"github.com/Nilesh-Coherent/baston-eventos/log"
	dbs "github.com/Nilesh-Coherent/common-service-evnt/pkg/db"
)

var DB *dbs.DB

func Initdb() {
	db, err := dbs.Initdb(config.Conf)
	if err != nil {
		log.Log.Println("Failed to Initiat DB Instance: ", err.Error())
		return
	}
	log.Log.Println("database insatnce created successfully.....")
	db.SetDebug()
	DB = db
	Migration()
}
