package service

import (
	emulator "github.com/duramash/constanta-emulator-task/pkg/models"
	"github.com/duramash/constanta-emulator-task/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(transaction emulator.TransactionModel) (int, error) {
	return s.repo.Create(transaction)
}
