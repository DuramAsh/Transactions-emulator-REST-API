package repository

import (
	"fmt"
	emulator "github.com/duramash/constanta-emulator-task/pkg/models"
	"github.com/duramash/constanta-emulator-task/pkg/utility"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"time"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) Create(body emulator.TransactionModel) error {

	currentTime := time.Now()
	status := "NEW"
	seed := rand.Intn(10)
	if seed == 7 {
		status = "ERROR"
	}
	query := "INSERT INTO transactions (user_id, user_email, amount, currency, date_of_creation, date_of_last_change, status) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.db.Exec(query, body.UserID, body.UserEmail, body.Amount, body.Currency, currentTime, currentTime, status)
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionPostgres) ChangeStatus(transactionId int, body emulator.TransactionModel) (string, error) {

	currentStatus, err := r.GetStatus(transactionId)
	if err != nil {
		return "", err
	}
	if utility.StatusIsTerminal(currentStatus) {
		return "", fmt.Errorf("status is terminal and cannot be modified")
	}
	if !utility.StatusIsAvailable(body.Status) {
		return "", fmt.Errorf("new status is incorrect")
	}
	if !r.TransactionExists(transactionId) {
		return "", fmt.Errorf("transaction DNE")
	}

	query := "UPDATE transactions SET status = $1, date_of_last_change = $2 WHERE id = $3"
	_, err = r.db.Exec(query, body.Status, time.Now(), transactionId)
	if err != nil {
		return "", err
	}
	return body.Status, nil
}

func (r *TransactionPostgres) GetStatus(transactionId int) (string, error) {
	var status string
	if r.TransactionExists(transactionId) {
		query := "SELECT status FROM transactions WHERE id = $1"
		err := r.db.Get(&status, query, transactionId)
		if err != nil {
			return "", err
		}
		return status, nil
	}

	return "", fmt.Errorf("transaction DNE")
}

func (r *TransactionPostgres) Cancel(transactionId int) error {
	if !r.TransactionExists(transactionId) {
		return fmt.Errorf("transaction DNE")
	}

	currentStatus, err := r.GetStatus(transactionId)
	if utility.StatusIsTerminal(currentStatus) {
		return fmt.Errorf("status is terminal, thus the current transaction cannot be deleted")
	}

	query := "DELETE FROM transactions WHERE id = $1"
	_, err = r.db.Exec(query, transactionId)
	if err != nil {
		return err
	}
	return nil

}

func (r *TransactionPostgres) GetByUserId(userId int) ([]emulator.TransactionModel, error) {
	var transactions []emulator.TransactionModel
	query := "SELECT * FROM transactions WHERE user_id = $1"
	err := r.db.Select(&transactions, query, userId)
	if err != nil {
		return nil, err
	}
	if len(transactions) == 0 {
		return nil, fmt.Errorf("no such user's transactions exist")
	}
	return transactions, nil
}

func (r *TransactionPostgres) GetByUserEmail(userEmail string) ([]emulator.TransactionModel, error) {
	var transactions []emulator.TransactionModel
	query := "SELECT * FROM transactions WHERE user_email = $1"
	err := r.db.Select(&transactions, query, userEmail)
	if err != nil {
		return nil, err
	}
	if len(transactions) == 0 {
		return nil, fmt.Errorf("no such user's transactions exist")
	}
	return transactions, nil
}
