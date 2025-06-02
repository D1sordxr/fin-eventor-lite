package dto

type DTO struct {
	ID      string  `json:"id,omitempty"`
	UserID  string  `json:"user_id,omitempty"`
	Balance float64 `json:"balance,omitempty"`
}
