package repository

import (
	emulator "github.com/duramash/constanta-emulator-task/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	Create(transaction emulator.TransactionModel) error
	ChangeStatus(transactionId int, transactions emulator.TransactionModel) error
	GetStatus(transactionId int) (string, error)
	Cancel(transactionId int) error
	GetByUserId(userId int) ([]emulator.TransactionModel, error)
	GetByUserEmail(userEmail string) ([]emulator.TransactionModel, error)
}

type Repository struct {
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransactionPostgres(db),
	}
}
