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
	productRepo := repositories.NewProductRepo(db)
	suggestionRepo := repositories.NewSuggestionRepository(db)
	// Service
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)
	suggestionService := services.NewSuggestionService(suggestionRepo)
	// Handler
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	suggestionHandler := handlers.NewSuggestionHandler(suggestionService)
	// Router et serveur (utilise Gin ici)
	router := server.SetupRouter(userHandler, productHandler, suggestionHandler)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erreur serveur:", err)
	}
}
