package emulator

import "time"

type TransactionModel struct {
	ID               int       `json:"id" db:"id"`
	UserID           int       `json:"user_id" db:"user_id"`
	UserEmail        string    `json:"user_email" db:"user_email"`
	Amount           int       `json:"amount" db:"amount"`
	Currency         string    `json:"currency" db:"currency"`
	DateOfCreation   time.Time `json:"date_of_creation" db:"date_of_creation"`
	DateOfLastChange time.Time `json:"date_of_last_change" db:"date_of_last_change"`
	Status           string    `json:"status" db:"status"`
}
