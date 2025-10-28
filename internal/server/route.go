package server

import (
	"net/http"
	"stockCorpo-api/internal/handlers"
)

// SetupRoutes configure toutes les routes de l'API
func SetupRoutes(userHandler *handlers.UserHandler) {
	// Route pour créer un utilisateur
	http.HandleFunc("/api/users", userHandler.CreateUser)
}
