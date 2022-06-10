package repository

type Authorisation interface {
}

type Transaction interface {
}

type Repository struct {
	Authorisation
	Transaction
}

func NewRepository() *Repository {
	return &Repository{}
}
