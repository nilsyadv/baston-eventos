package model

import (
	uuid "github.com/satori/go.uuid"

	"github.com/Nilesh-Coherent/common-service-evnt/pkg/model"
)

type Event struct {
	model.Trandb
	EventID          string    `json:"event_id"`
	EventName        string    `json:"event_name"`
	EventLocation    string    `json:"event_location"`
	EventDate        string    `json:"event_date"`
	Category         Category  `json:"event_category"`
	CategoryID       uuid.UUID `json:"-"`
	EventDescription string    `json:"event_description"`
	PaymentDetails   []Payment `json:"payment_details" gorm:"foreignKey:TransactionID"`
}
