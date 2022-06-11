package repository

import "github.com/jmoiron/sqlx"

type Authorisation interface {
}

type Transaction interface {
}

type Repository struct {
	Authorisation
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
