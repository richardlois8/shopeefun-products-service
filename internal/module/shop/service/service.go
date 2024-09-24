package service

import (
	"codebase-app/internal/module/shop/entity"
	"codebase-app/internal/module/shop/ports"
	"codebase-app/pkg/errmsg"
	"context"

	"github.com/rs/zerolog/log"
)

var _ ports.ShopService = &shopService{}

type shopService struct {
	repo ports.ShopRepository
}

func NewShopService(repo ports.ShopRepository) *shopService {
	return &shopService{
		repo: repo,
	}
}

func (s *shopService) CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error) {
	var res = new(entity.CreateShopResponse)
	isUser, err := s.repo.IsUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if !isUser {
		log.Warn().Any("payload", req).Msg("service: User is not shop owner")
		return res, errmsg.NewCustomErrors(403, errmsg.WithMessage("User not found"))
	}

	res, err = s.repo.CreateShop(ctx, req)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("service: Failed to create shop")
		return nil, err
	}

	return res, nil
}

func (s *shopService) GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error) {
	return s.repo.GetShop(ctx, req)
}

func (s *shopService) DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error {
	isUser, err := s.repo.IsUser(ctx, req.UserId)
	if err != nil {
		return err
	}

	if !isUser {
		log.Warn().Any("payload", req).Msg("service: User is not shop owner")
		return errmsg.NewCustomErrors(403, errmsg.WithMessage("User not found"))
	}

	return s.repo.DeleteShop(ctx, req)
}

func (s *shopService) UpdateShop(ctx context.Context, req *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error) {
	var res = new(entity.UpdateShopResponse)
	isUser, err := s.repo.IsUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if !isUser {
		log.Warn().Any("payload", req).Msg("service: User is not shop owner")
		return res, errmsg.NewCustomErrors(403, errmsg.WithMessage("User not found"))
	}

	res, err = s.repo.UpdateShop(ctx, req)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("service: Failed to update shop")
		return nil, err
	}

	return res, nil
}

func (s *shopService) GetShops(ctx context.Context, req *entity.ShopsRequest) (*entity.ShopsResponse, error) {
	return s.repo.GetShops(ctx, req)
}

func (s *shopService) IsUser(ctx context.Context, userId string) (bool, error) {
	return s.repo.IsUser(ctx, userId)
}