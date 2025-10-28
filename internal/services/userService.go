package services

import (
	"errors"
	"stockCorpo-api/internal/models"
	"stockCorpo-api/internal/repositories"

	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepository}
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

func (s *UserService) CreateUser(ctx context.Context, user *models.Users) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Pwd == "" {
		return errors.New("password is required")
	}
	// Hashing du mdp
	hashPwd, err := HashPassword(user.Pwd)
	if err != nil {
		return err
	}
	user.Pwd = hashPwd
	return s.userRepo.Create(ctx, user)
}
