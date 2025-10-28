package main

import (
	"log"
	"net/http"
	"stockCorpo-api/internal/database"
	"stockCorpo-api/internal/handlers"
	"stockCorpo-api/internal/repositories"
	"stockCorpo-api/internal/server"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	defer db.Close()

	// Initialisation du repository
	userRepo := repositories.NewUserRepo(db)

	// Initialisation du handler
	userHandler := handlers.NewUserHandler(userRepo)

	// Configuration des routes
	server.SetupRoutes(userHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Erreur serveur:", err)
	}
}
