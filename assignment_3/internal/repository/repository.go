package repository

import (
	"assignment_3/internal/repository/postgres"
)

type Repositories struct {
	UserRepo *postgres.UserRepository
}

func NewRepositories(db *postgres.Dialect) *Repositories {
	return &Repositories{
		UserRepo: postgres.NewUserRepository(db),
	}
}
