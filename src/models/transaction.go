package emulator

import "time"

type TransactionModel struct {
	ID               int       `json:"id" db:"id"`
	UserID           int       `json:"user_id" db:"user_id" binding:"required"`
	UserEmail        string    `json:"user_email" db:"user_email" binding:"required"`
	Amount           int       `json:"amount" db:"amount" binding:"required"`
	Currency         string    `json:"currency" db:"currency" binding:"required"`
	DateOfCreation   time.Time `json:"date_of_creation" db:"date_of_creation"`
	DateOfLastChange time.Time `json:"date_of_last_change" db:"date_of_last_change"`
	Status           string    `json:"status" db:"status"`
}

type Status struct {
	Status string `json:"status" db:"status" binding:"required"`
}
