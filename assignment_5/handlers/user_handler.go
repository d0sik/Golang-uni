package handlers

import (
	"assignment_5/repository"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 5
	}

	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	gender := r.URL.Query().Get("gender")
	orderBy := r.URL.Query().Get("order_by")

	res, err := h.repo.GetPaginatedUsers(page, pageSize, name, email, gender, orderBy)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) GetCommonFriends(w http.ResponseWriter, r *http.Request) {

	u1, _ := strconv.Atoi(r.URL.Query().Get("user1"))
	u2, _ := strconv.Atoi(r.URL.Query().Get("user2"))

	res, err := h.repo.GetCommonFriends(u1, u2)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)
}
