package service

import (
	emulator "github.com/duramash/constanta-emulator-task/pkg/models"
	"github.com/duramash/constanta-emulator-task/pkg/repository"
)

type Transaction interface {
	Create(transaction emulator.TransactionModel) (int, error) {

	}
}

type Service struct {
	Transaction
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
