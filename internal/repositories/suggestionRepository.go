package repositories

import (
	"context"
	"fmt"
	"stockCorpo-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// connection à la bdd
type SuggestionRepository struct {
	pool *pgxpool.Pool
}

// cree un objet SuggestionRepository (la poo est trop bizarre dans cette merde)
func NewSuggestionRepository(pool *pgxpool.Pool) *SuggestionRepository {
	return &SuggestionRepository{pool: pool}
}

// creer une suggestion
func (repository *SuggestionRepository) Create(ctx context.Context, suggestion *models.Suggestion) error {
	query := "INSERT INTO Suggestion(idUsers, suggestion) VALUES ($1, $2)"
	_, err := repository.pool.Exec(ctx, query, suggestion.IdUsers, suggestion.Suggestion)
	if err != nil {
		return fmt.Errorf("Echec de l'ajout de suggestion", err)
	}
	return nil
}

// supprimer une suggestion
func (repository *SuggestionRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM Suggestion WHERE idUsers = $1"
	_, err := repository.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Echec de la suppression de suggestion", err)
	}
	return nil
}

// avoir toute les suggestions
func (repository *SuggestionRepository) GetAll(ctx context.Context) ([]models.Suggestion, error) {
	var suggestions []models.Suggestion
	rows, err := repository.pool.Query(ctx, "SELECT suggestion FROM Suggestion")
	if err != nil {
		return nil, fmt.Errorf("echec de la recuperation des suggestions : %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var suggestion models.Suggestion
		if err := rows.Scan(&suggestion.Suggestion); err != nil {
			return nil, fmt.Errorf("echec du scan de la ligne : %w", err)
		}
		suggestions = append(suggestions, suggestion)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erreur de lecture des lignes: %w", err)
	}
	return suggestions, nil
}

func (repository *SuggestionRepository) Update(ctx context.Context, suggestion *models.Suggestion) error {
	query := `
		UPDATE Suggestion
		SET suggestion = $1
		WHERE idsuggestion = $2
	`

	_, err := repository.pool.Exec(ctx, query, suggestion.Suggestion, suggestion.IdSuggestion)
	if err != nil {
		return fmt.Errorf("échec de la mise à jour de la suggestion : %w", err)
	}
	return nil
}
