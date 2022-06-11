package repository

import (
	emulator "github.com/duramash/constanta-emulator-task/pkg/models"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) Create(transaction emulator.TransactionModel) (int, error) {
	query := "INSERT INTO transactions (user_id, user_email, amount, currency) VALUES ($1, $2, $3, $4) returning id"
}
