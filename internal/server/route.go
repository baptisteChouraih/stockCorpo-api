package server

import (
	"stockCorpo-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(userHandler *handlers.UserHandler) *gin.Engine {
	// Ton code ici
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/users", userHandler.CreateUser)
	}

	return router
}
