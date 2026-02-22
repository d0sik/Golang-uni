package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"assignment_3/internal/handler"
	"assignment_3/internal/middleware"
	"assignment_3/internal/repository"
	"assignment_3/internal/repository/postgres"
	"assignment_3/internal/usecase"
	"assignment_3/pkg/modules"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	ctx := context.Background()

	dbConfig := &modules.PostgreConfig{
		Host:        os.Getenv("DB_HOST"),
		Port:        os.Getenv("DB_PORT"),
		Username:    os.Getenv("DB_USER"),
		Password:    os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		SSLMode:     os.Getenv("DB_SSLMODE"),
		ExecTimeout: 5 * time.Second,
	}

	// Connecting to DB
	db := postgres.NewPGXDialect(ctx, dbConfig)

	// Layers
	repos := repository.NewRepositories(db)
	userUsecase := usecase.NewUserUsecase(repos.UserRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	mux := http.NewServeMux()

	// Healthcheck
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Routes
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

	// Middleware с API KEY из .env
	apiKey := os.Getenv("API_KEY")
	handlerWithMiddleware := middleware.Logging(
		middleware.APIKeyAuthWithKey(mux, apiKey),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handlerWithMiddleware))
}
