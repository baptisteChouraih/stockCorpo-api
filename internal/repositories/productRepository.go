package repositories

import (
	"context"
	"stockCorpo-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

// Create : permet de cree un produit dans la bdd
func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	query := "INSERT INTO Product (types, price, stock) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(ctx, query, product.Types, product.Price, product.Stock)
	return err
}

// GetId : permet de retrouver l'id d'un produit pour le modifier
func (r *ProductRepository) GetId(ctx context.Context, id int, product *models.Product) error {
	query := "SELECT idproduct FROM Product WHERE idproduct=$1"
	err := r.db.QueryRow(ctx, query, id).Scan(
		&product.IdProduct,
	)
	return err
}

// EditStock : permet de modifier le stock (true ou false)
func (r *ProductRepository) EditStock(ctx context.Context, product *models.Product) error {
	query := "UPDATE Product SET stock = $1 WHERE idproduct=$2"
	_, err := r.db.Exec(ctx, query, product.Stock, product.IdProduct)
	return err
}
