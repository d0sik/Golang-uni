package repo

import (
	"assignment_7/internal/entity"
	"assignment_7/pkg/postgres"
)

type UserRepo struct {
	DB *postgres.Postgres
}

func (r *UserRepo) Create(u *entity.User) error {
	return r.DB.Conn.Create(u).Error
}

func (r *UserRepo) GetByUsername(username string) (*entity.User, error) {
	var u entity.User
	err := r.DB.Conn.Where("username = ?", username).First(&u).Error
	return &u, err
}

func (r *UserRepo) GetByID(id string) (*entity.User, error) {
	var u entity.User
	err := r.DB.Conn.First(&u, "id = ?", id).Error
	return &u, err
}

func (r *UserRepo) Promote(id string) error {
	return r.DB.Conn.Model(&entity.User{}).
		Where("id = ?", id).
		Update("role", "admin").Error
}
