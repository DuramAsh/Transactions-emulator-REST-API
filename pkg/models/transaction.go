package emulator

import "time"

type Transaction struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	UserEmail        string    `json:"user_email"`
	Amount           int       `json:"amount"`
	Currency         string    `json:"currency"`
	DateOfCreation   time.Time `json:"date_of_creation"`
	DateOfLastChange time.Time `json:"date_of_last_change"`
	Status           string    `json:"status"`
}
