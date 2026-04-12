package usecase

import (
	"assignment_7/internal/entity"
	"assignment_7/internal/usecase/repo"
	"assignment_7/utils"
)

type UserUseCase struct {
	R *repo.UserRepo
}

func (u *UserUseCase) Register(user *entity.User) error {
	return u.R.Create(user)
}

func (u *UserUseCase) Login(dto *entity.LoginUserDTO) (string, error) {
	user, _ := u.R.GetByUsername(dto.Username)

	if !utils.CheckPassword(user.Password, dto.Password) {
		return "", nil
	}

	return utils.GenerateJWT(user.ID.String(), user.Role)
}

func (u *UserUseCase) GetMe(id string) (*entity.User, error) {
	return u.R.GetByID(id)
}

func (u *UserUseCase) Promote(id string) error {
	return u.R.Promote(id)
}
