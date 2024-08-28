package ports

import (
	"codebase-app/internal/module/product/entity"
	"context"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, shop *entity.CreateProductRequest) (*entity.CreateProductResponse, error)
	GetDetailProduct(ctx context.Context, shop *entity.GetProductDetailRequest) (*entity.GetProductDetailResponse, error)
	UpdateProduct(ctx context.Context, shop *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, shop *entity.DeleteProductRequest) (*entity.DeleteProductResponse, error)
	GetProducts(ctx context.Context, shop *entity.GetProductsRequest) (*entity.GetProductsResponse, error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, shop *entity.CreateProductRequest) (*entity.CreateProductResponse, error)
	GetDetailProduct(ctx context.Context, shop *entity.GetProductDetailRequest) (*entity.GetProductDetailResponse, error)
	UpdateProduct(ctx context.Context, shop *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, shop *entity.DeleteProductRequest) (*entity.DeleteProductResponse, error)
	GetProducts(ctx context.Context, shop *entity.GetProductsRequest) (*entity.GetProductsResponse, error)
}
