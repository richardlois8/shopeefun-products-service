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
		query = `INSERT INTO product (name, brand, price, stock, category_id, shop_id) VALUES (?, ?, ?, ?, ?, ?) RETURNING id`
	)

	err := r.db.QueryRowContext(ctx, r.db.Rebind(query),
		req.Name,
		req.Brand,
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

func (r *productRepository) GetDetailProduct(ctx context.Context, req *entity.GetProductDetailRequest) (*entity.GetProductDetailResponse, error) {
	var resp = new(entity.GetProductDetailResponse)
	var (
		query = `SELECT 
			p.id, 
			p.name, 
			p.price, 
			p.stock, 
			p.category_id,
			c.name as category_name, 
			COALESCE(p.description, '') as description, 
			COALESCE(p.image_url, '') as image_url,
			p.shop_id, 
			shops.name as shop_name,
			shops.description as shop_description
		FROM 
			product p 
		JOIN 
			category c 
		ON 
			p.category_id = c.id 
		JOIN 
			shops
		ON
			p.shop_id = shops.id
		WHERE 
			p.id = ?`
	)

	err := r.db.QueryRowxContext(
		ctx, r.db.Rebind(query), req.Id).Scan(
			&resp.Id,
			&resp.Name,
			&resp.Price,
			&resp.Stock,
			&resp.Category.Id,
			&resp.Category.Name,
			&resp.Description,
			&resp.ImageUrl,
			&resp.Shop.Id,
			&resp.Shop.Name,
			&resp.Shop.Description,
		)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetDetailProduct - Failed to get product detail")
		return nil, err
	}
	return resp, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error) {
	var resp = new(entity.UpdateProductResponse)
	var (
		query = `UPDATE product 
			SET name=?, 
				brand=?,
				price=?, 
				stock=?, 
				category_id=?, 
				description=?, 
				image_url=?, 
				updated_at = NOw() 
			WHERE id = ? AND shop_id=? 
			RETURNING id`
	)

	err := r.db.QueryRowContext(ctx, r.db.Rebind(query),
		req.Name,
		req.Brand,
		req.Price,
		req.Stock,
		req.CategoryId,
		req.Description,
		req.ImageUrl,
		req.Id,
		req.ShopId).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::UpdateProduct - Failed to update product")
		return nil, err
	}
	return resp, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) (*entity.DeleteProductResponse, error) {
	var resp = new(entity.DeleteProductResponse)
	var (
		query = `UPDATE product SET deleted_at = NOW() WHERE id = ? RETURNING id`
	)

	err := r.db.QueryRowContext(ctx, r.db.Rebind(query), req.Id).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::DeleteProduct - Failed to delete product")
		return nil, err
	}
	return resp, nil
}

func (r *productRepository) GetProducts(ctx context.Context, req *entity.GetProductsRequest) (*entity.GetProductsResponse, error) {
	type dao struct{
		TotalData int `db:"total_data"`
		entity.ProductItem
	}

	var(
		resp 	= new(entity.GetProductsResponse)
		data 	= make([]dao, 0, req.Paginate)
		query 	= `
			SELECT
				COUNT(p.id) OVER() as total_data,
				p.id,
				p.name,
				p.brand,
				p.price,
				p.stock,
				p.category_id,
				p.shop_id,
				COALESCE(p.description, '') as description,
				COALESCE(p.image_url, '') as image_url
			FROM
				product p
			WHERE
				deleted_at IS NULL
		`
		args []interface{}
	)

	if req.ProductName != "" {
		query += " AND p.name ILIKE ?"
		args = append(args, "%"+req.ProductName+"%")
		
	}
	if req.CategoryId != "" {
		query += " AND p.category_id = ?"
		args = append(args, req.CategoryId)
	}
	if req.Brand != "" {
		query += " AND p.brand ILIKE ?"
		args = append(args, "%"+req.Brand+"%")
	}
	if req.MinPrice > 0 {
		query += " AND p.price >= ?"
		args = append(args, req.MinPrice)
	}
	if req.MaxPrice > 0 {
		query += " AND p.price <= ?"
		args = append(args, req.MaxPrice)
	}

	// Pagination
	query += " LIMIT ? OFFSET ?"
	args = append(args, req.Paginate, (req.Page-1)*req.Paginate)

	err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), 
	args...)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetProducts - Failed to get products")
		return nil, err
	}

	if len(data) > 0 {
		resp.Meta.TotalData = data[0].TotalData
	}

	for _, d := range data {
		resp.Items = append(resp.Items, d.ProductItem)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}