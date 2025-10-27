package handlers

import (
	"encoding/json"
	"net/http"
	"stockCorpo-api/internal/models"
	"stockCorpo-api/internal/repositories"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

// NewUserHandler : cree une nouvelle instance de UserHandler tqt, je sais ce que je fais
func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo: repo}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Vérifie si la methode utiliser est POST (pour cree un user ta vu)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		//Si c est pas post return erreur jsp combien je sais juste ca marche pas
		return
	}

	var user models.Users
	// Decode le Json
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Crée l'utilisateur dans la base de données (pratique)
	err := handler.userRepo.Create(r.Context(), &user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	// Retourne une réponse de succès
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
