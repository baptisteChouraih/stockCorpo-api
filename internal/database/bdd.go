package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connection() (*pgxpool.Pool, error) {
	bdd := fmt.Sprintf(
		os.Getenv("DB_URL"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	pool, err := pgxpool.New(context.Background(), bdd)
	if err != nil {
		return nil, fmt.Errorf("echec de la connection a la base %w", err)
	}
	return pool, nil

}
