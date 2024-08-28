package ports

import (
	"codebase-app/internal/module/product/entity"
	"context"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, shop *entity.CreateProductRequest) (*entity.CreateProductResponse, error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, shop *entity.CreateProductRequest) (*entity.CreateProductResponse, error)
}
