package repositories

import (
	"context"
	"stockCorpo-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"

	"golang.org/x/crypto/bcrypt"
)

// UserRepository : connexion a la base de donnee
type UserRepository struct {
	db *pgxpool.Pool
}

// NewUserRepo : cree une sorte d objet UserRepo de ce que j ai compris
func NewUserRepo(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

// HashPassword : permet de hasher les mdp
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword : vérifie si un mot de passe correspond au hash pour le login tqt, je prends de l'avance grr grr gang
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Create : permet de cree un user dans la bdd
func (r *UserRepository) Create(ctx context.Context, user *models.Users) error {
	// Hashing du mdp
	hashPwd, err := HashPassword(user.Pwd)
	if err != nil {
		return err
	}

	// Insertion dans la bdd
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err = r.db.Exec(ctx, query, user.Name, user.Email, hashPwd)
	return err
}
