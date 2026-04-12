package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string
	Email    string
	Password string
	Role     string
}
