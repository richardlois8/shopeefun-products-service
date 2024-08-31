package entity

import "codebase-app/pkg/types"

type CreateProductRequest struct {
	// UserId string `validate:"validate"`

	Name        string  `json:"name" validate:"required,min=3,max=100" db:"name"`
	Price       float64 `json:"price" validate:"required" db:"price"`
	Stock       int     `json:"stock" validate:"required,min=1" db:"stock"`
	CategoryId  string  `json:"category_id" validate:"required,uuid" db:"category_id"`
	ShopId      string  `json:"shop_id" validate:"required,uuid" db:"shop_id"`
	Description string  `json:"description" db:"description"`
	ImageUrl    string  `json:"image_url" db:"image_url"`
}

type CreateProductResponse struct {
	Id string `json:"id" db:"id"`
}

type GetProductDetailRequest struct {
	Id string `validate:"uuid" db:"id"`
}

type GetProductDetailResponse struct {
	Id          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Price       float64 `json:"price" db:"price"`
	Stock       int     `json:"stock" db:"stock"`
	Category  CategoryItem  `json:"category"`
	Description *string  `json:"description" db:"description"`
	ImageUrl    *string  `json:"image_url" db:"image_url"`
	Shop ShopItem `json:"shop"`
}

type ShopItem struct{
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}	
type UpdateProductRequest struct {
	UserId string `prop:"user_id" validate:"uuid" db:"user_id"`

	Id          string  `params:"id" validate:"uuid" db:"id"`
	Name        string  `json:"name" validate:"required,min=3,max=100" db:"name"`
	Price       float64 `json:"price" validate:"required" db:"price"`
	Stock       int     `json:"stock" validate:"required,min=1" db:"stock"`
	CategoryId  string  `json:"category_id" validate:"required,uuid" db:"category_id"`
	ShopId      string  `json:"shop_id" validate:"required,uuid" db:"shop_id"`
	Description string  `json:"description" db:"description"`
	ImageUrl    string  `json:"image_url" db:"image_url"`
}

type UpdateProductResponse struct {
	Id string `json:"id" db:"id"`
}

type DeleteProductRequest struct {
	UserId string `prop:"user_id" validate:"uuid" db:"user_id"`

	Id string `validate:"uuid,required" db:"id"`
}

type DeleteProductResponse struct {
	Id string `json:"id" db:"id"`
}

type GetProductsRequest struct {
	UserId   string `prop:"user_id" validate:"uuid"`
	ProductName string `query:"name"`
	Brand string `query:"brand"`
	CategoryId string `query:"category"`
	MinPrice float64 `query:"min_price"`
	MaxPrice float64 `query:"max_price"`

	Page     int    `query:"page" validate:"required"`
	Paginate int    `query:"paginate" validate:"required"`
}

func (r *GetProductsRequest) SetDefault() {
	if r.Page < 1 {
		r.Page = 1
	}

	if r.Paginate < 1 {
		r.Paginate = 10
	}
}

type ProductItem struct {
	Id          string  `params:"id" validate:"uuid" db:"id"`
	Name        string  `json:"name" validate:"required,min=3,max=100" db:"name"`
	Brand	   	string  `json:"brand" db:"brand"`
	Price       float64 `json:"price" validate:"required" db:"price"`
	Stock       int     `json:"stock" validate:"required,min=1" db:"stock"`
	CategoryId  string  `json:"category_id" validate:"required,uuid" db:"category_id"`
	ShopId      string  `json:"shop_id" validate:"required,uuid" db:"shop_id"`
	Description string  `json:"description" db:"description"`
	ImageUrl    string  `json:"image_url" db:"image_url"`
}

type GetProductsResponse struct {
	Items []ProductItem `json:"items"`
	Meta  types.Meta    `json:"meta"`
}

type CategoryItem struct{
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
