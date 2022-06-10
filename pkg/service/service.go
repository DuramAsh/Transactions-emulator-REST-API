package service

import "github.com/duramash/constanta-emulator-task/pkg/repository"

type Authorisation interface {
}

type Transaction interface {
}

type Service struct {
	Authorisation
	Transaction
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
