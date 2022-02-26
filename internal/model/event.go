package model

import (
	uuid "github.com/satori/go.uuid"

	"github.com/Nilesh-Coherent/common-service-evnt/pkg/model"
)

type Event struct {
	model.Trandb
	EventName        string    `json:"event_name"`
	EventLocation    string    `json:"event_location"`
	EventDate        string    `json:"event_date"`
	Category         *Category `json:"-" gorm:"foreignKey:ID"`
	CategoryID       uuid.UUID `json:"category_id"`
	EventDescription string    `json:"event_description"`
	PaymentDetails   *Payment  `json:"payment_details,omitempty" gorm:"foreignKey:EventID"`
}
