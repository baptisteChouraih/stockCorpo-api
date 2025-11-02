package handlers

import (
	"net/http"
	"stockCorpo-api/internal/models"
	"stockCorpo-api/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SuggestionHandler struct {
	service *services.SuggestionService
}

func NewSuggestionHandler(service *services.SuggestionService) *SuggestionHandler {
	return &SuggestionHandler{service: service}
}

func (handler *SuggestionHandler) CreateSuggestion(c *gin.Context) {
	var suggestion models.Suggestion

	if err := c.ShouldBindJSON(&suggestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json invalide"})
		return
	}

	if err := handler.service.Create(c.Request.Context(), &suggestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "la suggestion a été créé avec succès", "suggestion": suggestion})
}

func (handler *SuggestionHandler) GetAllSuggestions(c *gin.Context) {
	suggestions, err := handler.service.GetSuggestions(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"suggestions": suggestions})
}

func (handler *SuggestionHandler) DeleteSuggestion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalide"})
		return
	}

	if err := handler.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "suggestion supprimée"})

}

func (handler *SuggestionHandler) UpdateSuggestion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalide"})
		return
	}

	var suggestion models.Suggestion
	if err := c.ShouldBindJSON(&suggestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json invalide"})
		return
	}

	suggestion.IdSuggestion = id

	if err := handler.service.Update(c.Request.Context(), &suggestion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "suggestion mise à jour"})
}
