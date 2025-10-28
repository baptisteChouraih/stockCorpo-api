package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Connection() (*pgxpool.Pool, error) {
	// Charge le fichier .env
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("erreur chargement .env: %w", err)
	}

	// Récupère la connection string
	connString := os.Getenv("DB_URL")
	if connString == "" {
		return nil, fmt.Errorf("DB_URL n'est pas définie")
	}

	// Connexion
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("echec connexion: %w", err)
	}

	return pool, nil
}
