package main

import (
	"log"
	"stockCorpo-api/internal/database"
	"stockCorpo-api/internal/handlers"
	"stockCorpo-api/internal/repositories"
	"stockCorpo-api/internal/server"
)

func main() {
	// Connexion DB
	db, err := database.Connection()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	defer db.Close()
	// Repository
	userRepo := repositories.NewUserRepo(db)
	// Handler
	userHandler := handlers.NewUserHandler(userRepo)
	// Router et serveur (utilise Gin ici)
	router := server.SetupRoutes(userHandler)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erreur serveur:", err)
	}
}
