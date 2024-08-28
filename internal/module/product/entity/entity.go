package entity

type CreateProductRequest struct {
	UserId string `query:"user_id" validate:"required,uuid"`

	Name string `json:"name" validate:"required,min=3,max=100" db:"name"`
	Price float64 `json:"price" validate:"required" db:"price"`
	Stock int `json:"stock" validate:"required,min=1" db:"stock"`
	CategoryId string `json:"category_id" validate:"required,uuid" db:"category_id"`
	ShopId string `json:"shop_id" validate:"required,uuid" db:"shop_id"`
	Description string `json:"description" db:"description"`
	ImageUrl string `json:"image_url" db:"image_url"`
}

type CreateProductResponse struct {
	Id string `json:"id" db:"id"`
}

type ProductResult struct {
}
