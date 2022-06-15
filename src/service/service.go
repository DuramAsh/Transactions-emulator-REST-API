package service

import (
	emulator "github.com/duramash/constanta-emulator-task/src/models"
	"github.com/duramash/constanta-emulator-task/src/repository"
)

type Transaction interface {
	Create(body emulator.TransactionModel) error
	ChangeStatus(transactionId int, body emulator.Status) (string, error)
	GetStatus(transactionId int) (string, error)
	Cancel(transactionId int) error
	GetByUserId(userId int) ([]emulator.TransactionModel, error)
	GetByUserEmail(userEmail string) ([]emulator.TransactionModel, error)
}

type Service struct {
	Transaction
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Transaction: NewTransactionService(repo.Transaction),
	}
}
