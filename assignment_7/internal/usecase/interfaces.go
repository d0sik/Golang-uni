package usecase

import "assignment_7/internal/entity"

type UserInterface interface {
	Register(*entity.User) error
	Login(*entity.LoginUserDTO) (string, error)
	GetMe(string) (*entity.User, error)
	Promote(string) error
}
