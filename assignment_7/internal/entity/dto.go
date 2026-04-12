package entity

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
