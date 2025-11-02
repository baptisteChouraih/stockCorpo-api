package services

import (
	"context"
	"fmt"
	"stockCorpo-api/internal/models"
	"stockCorpo-api/internal/repositories"
)

type SuggestionService struct {
	repository repositories.SuggestionRepository
}

// creer un objet SuggestionService
func NewSuggestionService(repository repositories.SuggestionRepository) *SuggestionService {
	return &SuggestionService{repository: repository}
}

func (service *SuggestionService) GetSuggestions(ctx context.Context) ([]models.Suggestion, error) {
	suggestions, err := service.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des suggestions : w%", err)
	}
	return suggestions, nil
}

func (service *SuggestionService) Create(ctx context.Context, suggestion *models.Suggestion) error {
	if suggestion.Suggestion == "" {
		return fmt.Errorf("le champ suggestion est vide")
	}
	return service.repository.Create(ctx, suggestion)
}

func (service *SuggestionService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("id invalide")
	}
	return service.repository.Delete(ctx, id)
}

func (service *SuggestionService) Update(ctx context.Context, suggestion *models.Suggestion) error {
	if suggestion.IdSuggestion <= 0 {
		return fmt.Errorf("id invalide")
	}
	if suggestion.Suggestion == "" {
		return fmt.Errorf("suggestion ne peut pas etre vide")
	}
	return service.repository.Update(ctx, suggestion)
}
