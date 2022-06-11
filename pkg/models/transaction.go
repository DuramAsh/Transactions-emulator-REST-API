package emulator

import "time"

type TransactionModel struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id" binding:"required"`
	UserEmail        string    `json:"user_email" binding:"required"`
	Amount           int       `json:"amount" binding:"required"`
	Currency         string    `json:"currency" binding:"required"`
	DateOfCreation   time.Time `json:"date_of_creation"`
	DateOfLastChange time.Time `json:"date_of_last_change"`
	Status           string    `json:"status"`
}
