package service

import (
	emulator "github.com/duramash/constanta-emulator-task/src/models"
	"github.com/duramash/constanta-emulator-task/src/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(body emulator.TransactionModel) error {
	return s.repo.Create(body)
}

func (s *TransactionService) ChangeStatus(transactionId int, body emulator.TransactionModel) (string, error) {
	return s.repo.ChangeStatus(transactionId, body)
}

func (s *TransactionService) GetStatus(transactionId int) (string, error) {
	return s.repo.GetStatus(transactionId)
}

func (s *TransactionService) Cancel(transactionId int) error {
	return s.repo.Cancel(transactionId)
}

func (s *TransactionService) GetByUserId(userId int) ([]emulator.TransactionModel, error) {
	return s.repo.GetByUserId(userId)
}

func (s *TransactionService) GetByUserEmail(userEmail string) ([]emulator.TransactionModel, error) {
	return s.repo.GetByUserEmail(userEmail)
}
