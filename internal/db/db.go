package db

import (
	"github.com/lin-sel/baston-eventos/config"
	"github.com/lin-sel/baston-eventos/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Createdb() {
	db, err := gorm.Open(mysql.Open(CreateDBString()), &gorm.Config{})
	if err != nil {
		log.Log.Println("Error in db Connection start:", err.Error())
		return
	}
	DB = db
}

func CreateDBString() string {
	return config.GetString("db_user") + ":" + config.GetString("db_pass") +
		"@tcp(" + config.GetString("db_host") + ":" + config.GetString("db_port") + ")" +
		"/" + config.GetString("db_name") + "?charset=utf8mb4&parseTime=True&loc=Local"
}
