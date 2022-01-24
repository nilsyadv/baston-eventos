package model

import (
	uuid "github.com/satori/go.uuid"

	"github.com/Nilesh-Coherent/common-service-evnt/pkg/model"
)

type Payment struct {
	model.Trandb
	TotalAmount       float64   `json:"total_amount"`
	PaidAmount        float64   `json:"paid_amount"`
	PendingAmount     float64   `json:"pending_amount"`
	TransactionRef    string    `json:"transaction_ref"`
	TransactionDate   string    `json:"transaction_date"`
	TransactionAmount string    `json:"transaction_amount"`
	EventID           uuid.UUID `json:"event_id"`
}
