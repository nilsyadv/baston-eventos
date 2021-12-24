package db

import (
	"github.com/Nilesh-Coherent/baston-eventos/internal/model"
	"github.com/Nilesh-Coherent/baston-eventos/log"
)

func Migration() {
	tables := []interface{}{
		&model.Category{},
		&model.Event{},
		&model.Payment{},
		&model.PaymentHistory{},
	}

	// automigrating all tables
	errs := DB.AutoMigrates(tables)
	if len(errs) > 0 {
		for _, err := range errs {
			log.Log.Fatal(err.Error())
		}
		return
	}

	log.Log.Println("db migration successfully done.")
}
