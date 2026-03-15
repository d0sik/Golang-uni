package repository

import (
	"assignment_5/models"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetPaginatedUsers(
	page int,
	pageSize int,
	name string,
	email string,
	gender string,
	orderBy string,
) (models.PaginatedResponse, error) {

	offset := (page - 1) * pageSize

	query := "SELECT id,name,email,gender,birth_date FROM users WHERE 1=1"

	args := []interface{}{}
	argID := 1

	if name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argID)
		args = append(args, "%"+name+"%")
		argID++
	}

	if email != "" {
		query += fmt.Sprintf(" AND email ILIKE $%d", argID)
		args = append(args, "%"+email+"%")
		argID++
	}

	if gender != "" {
		query += fmt.Sprintf(" AND gender=$%d", argID)
		args = append(args, gender)
		argID++
	}

	if orderBy == "" {
		orderBy = "id"
	}

	query += " ORDER BY " + orderBy
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argID, argID+1)

	args = append(args, pageSize, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return models.PaginatedResponse{}, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var u models.User

		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Gender, &u.BirthDate)

		users = append(users, u)
	}

	var total int

	r.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&total)

	return models.PaginatedResponse{
		Data:       users,
		TotalCount: total,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

func (r *UserRepository) GetCommonFriends(user1 int, user2 int) ([]models.User, error) {

	query := `
	SELECT u.id,u.name,u.email,u.gender,u.birth_date
	FROM users u
	JOIN user_friends f1 ON u.id=f1.friend_id
	JOIN user_friends f2 ON u.id=f2.friend_id
	WHERE f1.user_id=$1 AND f2.user_id=$2
	`

	rows, err := r.db.Query(query, user1, user2)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var u models.User

		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Gender, &u.BirthDate)

		users = append(users, u)
	}

	return users, nil
}
