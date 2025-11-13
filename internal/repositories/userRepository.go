package repositories

import (
	"context"
	"stockCorpo-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// UserRepository : connexion a la base de donnee
type UserRepository struct {
	db *pgxpool.Pool
}

// NewUserRepo : cree une sorte d objet UserRepo de ce que j ai compris
func NewUserRepo(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

// Create : permet de cree un user dans la bdd
func (r *UserRepository) Create(ctx context.Context, user *models.Users) error {
	// Insertion dans la bdd
	query := "INSERT INTO users (name, email, pwd) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(ctx, query, user.Name, user.Email, user.Pwd)
	return err
}

// GetByEmail : trouver un utilisateur avec son email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.Users, error) {
	var user models.Users
	query := "SELECT idusers, name, email, pwd FROM users WHERE email = $1"
	err := r.db.QueryRow(ctx, query, email).Scan(&user.IdUsers, &user.Name, &user.Email, &user.Pwd)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
