package service

import (
	"codebase-app/internal/module/product/entity"
	"codebase-app/internal/module/product/ports"
	"context"
)

var _ ports.ProductService = &productService{}

type productService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *productService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error) {
	return s.repo.CreateProduct(ctx, req)
}

func (s *productService) GetDetailProduct(ctx context.Context, req *entity.GetProductDetailRequest) (*entity.GetProductDetailResponse, error) {
	return s.repo.GetDetailProduct(ctx, req)
}

func (s *productService) UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error) {
	return s.repo.UpdateProduct(ctx, req)
}

func (s *productService) DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) (*entity.DeleteProductResponse, error) {
	return s.repo.DeleteProduct(ctx, req)
}

func (s *productService) GetProducts(ctx context.Context, req *entity.GetProductsRequest) (*entity.GetProductsResponse, error) {
	return s.repo.GetProducts(ctx, req)
}