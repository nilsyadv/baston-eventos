package model

import (
	uuid "github.com/satori/go.uuid"

	"github.com/Nilesh-Coherent/common-service-evnt/pkg/model"
)

type Payment struct {
	model.Trandb
	TotalAmount    float64          `json:"total_amount"`
	PaidAmount     float64          `json:"paid_amount"`
	PendingAmount  float64          `json:"pending_amount"`
	PaymentHistory []PaymentHistory `json:"payment_history"`
	TransactionID  uuid.UUID        `json:"transaction_id"`
}

type PaymentHistory struct {
	model.Trandb
	TransactionID     string    `json:"transaction_id"`
	TransactionRef    string    `json:"transaction_ref"`
	TransactionDate   string    `json:"transaction_date"`
	TransactionAmount string    `json:"transaction_amount"`
	PaymentID         uuid.UUID `json:"payment_id"`
}
