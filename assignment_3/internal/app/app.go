package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"assignment_3/internal/handler"
	"assignment_3/internal/middleware"
	"assignment_3/internal/repository"
	"assignment_3/internal/repository/postgres"
	"assignment_3/internal/usecase"
	"assignment_3/pkg/modules"
)

func Run() {
	ctx := context.Background()

	dbConfig := &modules.PostgreConfig{
		Host:        "localhost",
		Port:        "5432",
		Username:    "postgres",
		Password:    "dosik13094664",
		DBName:      "golang_sis3",
		SSLMode:     "disable",
		ExecTimeout: 5 * time.Second,
	}

	// DB
	db := postgres.NewPGXDialect(ctx, dbConfig)

	// Layers
	repos := repository.NewRepositories(db)
	userUsecase := usecase.NewUserUsecase(repos.UserRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// Router
	mux := http.NewServeMux()

	// Healthcheck
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Users routes
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUserByID(w, r)
		case http.MethodPut:
			userHandler.UpdateUser(w, r)
		case http.MethodDelete:
			userHandler.DeleteUser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Middleware
	handlerWithMiddleware := middleware.Logging(
		middleware.APIKeyAuth(mux),
	)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithMiddleware))
}
