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

func (r *TransactionPostgres) Create(transaction emulator.TransactionModel) error {

	query := "INSERT INTO transactions (user_id, user_email, amount, currency) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, transaction.UserID, transaction.UserEmail, transaction.Amount, transaction.Currency)
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionPostgres) ChangeStatus(transactionId int, transaction emulator.TransactionModel) error {
	query := "UPDATE transactions SET status = $1 WHERE id = $2"
	_, err := r.db.Exec(query, transaction.Status, transactionId)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionPostgres) GetStatus(transactionId int) (string, error) {
	var status string
	query := "SELECT status FROM transactions WHERE id = $1"
	err := r.db.Get(&status, query, transactionId)
	if err != nil {
		return "", err
	}
	return status, nil
}

func (r *TransactionPostgres) Cancel(transactionId int) error {
	query := "DELETE FROM transactions WHERE id = $1"
	_, err := r.db.Exec(query, transactionId)
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionPostgres) GetByUserId(userId int) ([]emulator.TransactionModel, error) {
	var transactions []emulator.TransactionModel
	query := "SELECT id, user_id, user_email, amount, currency FROM transactions WHERE user_id = $1"
	err := r.db.Select(&transactions, query, userId)
	if err != nil {
		return []emulator.TransactionModel{}, err
	}
	return transactions, nil
}

func (r *TransactionPostgres) GetByUserEmail(userEmail string) ([]emulator.TransactionModel, error) {
	var transactions []emulator.TransactionModel
	query := "SELECT id, user_id, user_email, amount, currency FROM transactions WHERE user_email = $1"
	err := r.db.Select(&transactions, query, userEmail)
	if err != nil {
		return []emulator.TransactionModel{}, err
	}
	return transactions, nil
}
