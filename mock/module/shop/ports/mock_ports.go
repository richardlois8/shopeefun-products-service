package mock_ports

import (
	"codebase-app/internal/module/shop/entity"
	"codebase-app/internal/module/shop/ports"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockShopRepo struct {
	mock.Mock
}

func NewMockShopRepo() *MockShopRepo {
	return &MockShopRepo{}
}

var _ ports.ShopRepository = &MockShopRepo{}

func (m *MockShopRepo) CreateShop(ctx context.Context, shop *entity.CreateShopRequest) (*entity.CreateShopResponse, error) {
	args := m.Called(ctx, shop)
	var (
		resp entity.CreateShopResponse
		err  error
	)

	if n, ok := args.Get(0).(entity.CreateShopResponse); ok {
		resp = n
	}
	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return &resp, err
}

func (m *MockShopRepo) GetShop(ctx context.Context, shop *entity.GetShopRequest) (*entity.GetShopResponse, error) {
	args := m.Called(ctx, shop)
	var (
		resp entity.GetShopResponse
		err  error
	)

	if n, ok := args.Get(0).(entity.GetShopResponse); ok {
		resp = n
	}
	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return &resp, err
}

func (m *MockShopRepo) UpdateShop(ctx context.Context, shop *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error) {
	args := m.Called(ctx, shop)
	var (
		resp entity.UpdateShopResponse
		err  error
	)

	if n, ok := args.Get(0).(entity.UpdateShopResponse); ok {
		resp = n
	}
	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return &resp, err
}

func (m *MockShopRepo) DeleteShop(ctx context.Context, shop *entity.DeleteShopRequest) error {
	args := m.Called(ctx, shop)
	var err error

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}

func (m *MockShopRepo) GetShops(ctx context.Context, shop *entity.ShopsRequest) (*entity.ShopsResponse, error) {
	args := m.Called(ctx, shop)
	var (
		resp entity.ShopsResponse
		err  error
	)

	if n, ok := args.Get(0).(entity.ShopsResponse); ok {
		resp = n
	}
	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return &resp, err
}

func (m *MockShopRepo) IsUser(ctx context.Context, userId string) (bool, error) {
	args := m.Called(ctx, userId)
	var (
		resp bool
		err  error
	)

	if n, ok := args.Get(0).(bool); ok {
		resp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return resp, err
}