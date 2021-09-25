package model

type Payment struct {
	TotalAmount    float64          `json:"total_amount"`
	PaidAmount     float64          `json:"paid_amount"`
	PendingAmount  float64          `json:"pending_amount"`
	PaymentHistory []PaymentHistory `json:"payment_history"`
}

type PaymentHistory struct {
	TransactionID     string `json:"transaction_id"`
	TransactionRef    string `json:"transaction_ref"`
	TransactionDate   string `json:"transaction_date"`
	TransactionAmount string `json:"transaction_amount"`
}
