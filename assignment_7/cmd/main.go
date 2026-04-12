package main

import (
	"log"
	"os"

	"assignment_7/internal/controller/http/v1"
	"assignment_7/internal/entity"
	"assignment_7/internal/usecase"
	"assignment_7/internal/usecase/repo"
	"assignment_7/pkg/postgres"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.New(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	db.Conn.AutoMigrate(&entity.User{})

	userRepo := &repo.UserRepo{DB: db}
	userUse := &usecase.UserUseCase{R: userRepo}

	r := gin.Default()

	v1.New(r, userUse)

	r.Run(":8080")
}
