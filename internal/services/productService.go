package services

import (
	"errors"
	"stockCorpo-api/internal/models"
	"stockCorpo-api/internal/repositories"

	"context"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepository}
}

func (p *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	if product.Types == "" {
		return errors.New("types is required")
	}

	if product.Price == nil {
		return errors.New("price is required")
	}

	if *product.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	if product.Stock == nil {
		return errors.New("stock is required")
	}

	return p.productRepo.Create(ctx, product)
}

func (p *ProductService) EditProduct(ctx context.Context, product *models.Product) error {
	if product.Stock == nil {
		return errors.New("stock is required")
	}
	return p.productRepo.EditStock(ctx, product)
}
