package main

import (
	"log"
	"stockCorpo-api/internal/database"
	"stockCorpo-api/internal/handlers"
	"stockCorpo-api/internal/repositories"
	"stockCorpo-api/internal/server"
	"stockCorpo-api/internal/services"
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
	// Service
	userService := services.NewUserService(userRepo)
	// Handler
	userHandler := handlers.NewUserHandler(userService)
	// Router et serveur (utilise Gin ici)
	router := server.SetupRoutes(userHandler)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erreur serveur:", err)
	}
}
