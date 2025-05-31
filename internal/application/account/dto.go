package account

type DTO struct {
	ID      string  `json:"id,omitempty"`
	UserID  string  `json:"user_id,omitempty"`
	Balance float64 `json:"balance,omitempty"`
}

type EventDTO struct {
	ID        string  `json:"id"`
	AccountID string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
}
