package repository

import (
	"codebase-app/internal/module/product/entity"
	"codebase-app/internal/module/product/ports"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var _ ports.ProductRepository = &productRepository{}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error) {
	var resp = new(entity.CreateProductResponse)
	var (
		query = `INSERT INTO product (name, price, stock, category_id, shop_id) VALUES (?, ?, ?, ?, ?) RETURNING id`
	)

	err := r.db.QueryRowContext(ctx, r.db.Rebind(query),
		req.Name,
		req.Price,
		req.Stock,
		req.CategoryId,
		req.ShopId).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::CreateProduct - Failed to create product")
		return nil, err
	}
	return resp, nil
}
