package main

import (
	"log"
	"net/http"

	"assignment_5/db"
	"assignment_5/handlers"
	"assignment_5/repository"
)

func main() {

	database := db.InitDB()

	repo := repository.NewUserRepository(database)

	handler := handlers.NewUserHandler(repo)

	http.HandleFunc("/users", handler.GetUsers)
	http.HandleFunc("/common_friends", handler.GetCommonFriends)

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}
