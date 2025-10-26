package handlers

import (
	"net/http"
	"stockCorpo-api/internal/repositories"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

// NewUserHandler : cree une nouvelle instance de UserHandler tqt, je sais ce que je fais
func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo: repo}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {}

//TODO : faire la fonction CreateUser
