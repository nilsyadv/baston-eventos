package model

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

	custerror "github.com/nilsyadv/baston-eventos/error"
	"github.com/nilsyadv/baston-eventos/internal/util"
	"github.com/nilsyadv/common-service-evnt/pkg/model"
)

type PaymentHistory struct {
	model.Trandb
	TransactionRef    string     `json:"transaction_ref"`
	TransactionDate   *time.Time `json:"transaction_date"`
	TransactionAmount float64    `json:"transaction_amount"`
	PaymentID         uuid.UUID  `json:"payment_id"`
}

func (hist *PaymentHistory) ValidatePayHist() *custerror.CustomeError {
	if hist.TransactionAmount <= 0 {
		er := custerror.CreateCustomeError("transaction amount can not be zero or less", nil,
			http.StatusBadRequest)
		return &er
	}
	if hist.PaymentID == uuid.Nil {
		er := custerror.CreateCustomeError("payment id can not be empty", nil,
			http.StatusBadRequest)
		return &er
	}

	// if id is null then create new id & save it
	if hist.ID == uuid.Nil {
		_, hist.ID = util.CreateID()
	}

	// if transaction date is null then assign new transaction date
	if hist.TransactionDate == nil {
		tms := time.Now()
		hist.TransactionDate = &tms
	}

	return nil
}
