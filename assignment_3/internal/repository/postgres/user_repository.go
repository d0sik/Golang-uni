package postgres

import (
	"assignment_3/pkg/modules"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *Dialect
}

func NewUserRepository(db *Dialect) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]modules.User, error) {
	var users []modules.User
	err := r.db.DB.Select(&users, "SELECT id, name, email, age FROM users")
	return users, err
}

func (r *UserRepository) GetUserByID(id int) (*modules.User, error) {
	var user modules.User
	err := r.db.DB.Get(&user, "SELECT id, name, email, age FROM users WHERE id=$1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(u modules.User) (int, error) {
	var id int
	query := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.DB.QueryRow(query, u.Name, u.Email, u.Age).Scan(&id)
	return id, err
}

func (r *UserRepository) UpdateUser(id int, u modules.User) error {
	result, err := r.db.DB.Exec(
		"UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4",
		u.Name, u.Email, u.Age, id,
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) (int64, error) {
	result, err := r.db.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
