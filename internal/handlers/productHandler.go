package handlers

import (
	"net/http"
	"stockCorpo-api/internal/models"
	"stockCorpo-api/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *services.ProductService
}

// NewProductHandler : cree une nouvelle instance de ProductHandler tqt, je sais ce que je fais
func NewProductHandler(serv *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: serv}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.productService.CreateProduct(c.Request.Context(), &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "Product created successfully!"})
}

func (h *ProductHandler) EditProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.productService.EditProduct(c.Request.Context(), &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "Product updated successfully!"})
}
