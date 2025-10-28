package handlers

import (
	"net/http"
	"stockCorpo-api/internal/models"
	"stockCorpo-api/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler : cree une nouvelle instance de UserHandler tqt, je sais ce que je fais
func NewUserHandler(serv *services.UserService) *UserHandler {
	return &UserHandler{userService: serv}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "User created successfully!"})
}
