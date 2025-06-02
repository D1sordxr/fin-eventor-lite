package dto

type EventDTO struct {
	ID        string  `json:"id"`
	AccountID string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
}
