package model

type Event struct {
	EventID          string    `json:"event_id"`
	EventName        string    `json:"event_name"`
	EventLocation    string    `json:"event_location"`
	EventDate        string    `json:"event_date"`
	EventCategory    ECategory `json:"event_category"`
	EventDescription string    `json:"event_description"`
	PaymentDetails   []Payment `json:"payment_details"`
}
