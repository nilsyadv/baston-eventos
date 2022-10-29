package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/nilsyadv/baston-eventos/internal/util"
	"github.com/nilsyadv/common-service-evnt/pkg/model"
)

type Payment struct {
	model.Trandb
	TotalAmount    float64          `json:"total_amount"`
	PaidAmount     float64          `json:"paid_amount"`
	PendingAmount  float64          `json:"pending_amount"`
	LastPaymentOn  *time.Time       `json:"last_payment"`
	PaymentHistory []PaymentHistory `json:"payment_histories,omitempty" gorm:"foreignKey:PaymentID;references:ID"`
	EventID        uuid.UUID        `json:"event_id"`
}

func (payment *Payment) PaymentCalculation() error {
	if payment.ID == uuid.Nil {
		_, payment.ID = util.CreateID()
	}
	for index, hist := range payment.PaymentHistory {
		payment.PaidAmount += hist.TransactionAmount

		// validate payment history
		err := hist.ValidatePayHist()
		if err != nil {
			return err
		}

		// assign transaction if paydate is null
		if payment.LastPaymentOn == nil {
			payment.LastPaymentOn = hist.TransactionDate
		}

		// getting latest transaction date
		if hist.TransactionDate.After(*payment.LastPaymentOn) {
			payment.LastPaymentOn = hist.TransactionDate
		}

		payment.PaymentHistory[index] = hist
	}

	payment.PendingAmount = payment.TotalAmount - payment.PaidAmount
	return nil
}
