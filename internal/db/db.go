package db

import (
	"github.com/nilsyadv/baston-eventos/config"
	"github.com/nilsyadv/baston-eventos/log"
	dbs "github.com/nilsyadv/common-service-evnt/pkg/db"
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
