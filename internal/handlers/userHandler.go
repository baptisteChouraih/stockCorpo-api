package handlers

import (
	"net/http"
	"stockCorpo-api/internal/models"
	"stockCorpo-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

// NewUserHandler : cree une nouvelle instance de UserHandler tqt, je sais ce que je fais
func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo: repo}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userRepo.Create(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "User created successfully!"})
}
