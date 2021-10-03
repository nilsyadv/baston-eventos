package db

import (
	"fmt"

	"github.com/lin-sel/baston-eventos/config"
	"github.com/lin-sel/baston-eventos/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initdb() {
	db := Createdb()
	if db == nil {
		log.Log.Println("Failed to Initiat DB Instance..")
		return
	}
	log.Log.Println("database insatnce created successfully.....")
	DB = db
}

func Createdb() *gorm.DB {
	dbconn := CreateDBString()
	fmt.Println("db Connection String:", dbconn)
	db, err := gorm.Open(postgres.Open(dbconn), &gorm.Config{})
	if err != nil {
		log.Log.Println("Error in db Connection start:", err.Error())
		log.Log.Fatal("Failed to connect db....")
		return nil
	}
	return db
}

func CreateDBString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.GetString("db.db_host"), config.GetString("db.db_user"), config.GetString("db.db_pass"), config.GetString("db.db_name"), config.GetString("db.db_port"), "disable", "Asia/Shanghai")
	// return config.GetString("db.db_user") + ":" + config.GetString("db.db_pass") +
	// 	"@tcp(" + config.GetString("db.db_host") + ":" + config.GetString("db.db_port") + ")" +
	// 	"/" + config.GetString("db.db_name") + "?charset=utf8mb4&parseTime=True&loc=Local"
}
