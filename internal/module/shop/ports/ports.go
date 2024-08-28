package ports

import (
	"codebase-app/internal/module/shop/entity"
	"context"
)

type ShopRepository interface {
	CreateShop(ctx context.Context, shop *entity.CreateShopRequest) (*entity.CreateShopResponse, error)
	GetShop(ctx context.Context, shop *entity.GetShopRequest) (*entity.GetShopResponse, error)
	UpdateShop(ctx context.Context, shop *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error)
	DeleteShop(ctx context.Context, shop *entity.DeleteShopRequest) error
	GetShops(ctx context.Context, shop *entity.ShopsRequest) (*entity.ShopsResponse, error)
}

type ShopService interface {
		CreateShop(ctx context.Context, shop *entity.CreateShopRequest) (*entity.CreateShopResponse, error)
	GetShop(ctx context.Context, shop *entity.GetShopRequest) (*entity.GetShopResponse, error)
	UpdateShop(ctx context.Context, shop *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error)
	DeleteShop(ctx context.Context, shop *entity.DeleteShopRequest) error
	GetShops(ctx context.Context, shop *entity.ShopsRequest) (*entity.ShopsResponse, error)
}
