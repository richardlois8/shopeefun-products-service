package service

import (
	"codebase-app/internal/module/shop/entity"
	"codebase-app/internal/module/shop/ports"
	mockPort "codebase-app/mock/module/shop/ports"
	"codebase-app/pkg/errmsg"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockShopService struct {
	mock.Mock
}

func NewMockShopService() *MockShopService {
	return &MockShopService{}
}

type ServiceList struct{
	suite.Suite
	mockShopRepo *mockPort.MockShopRepo
	service ports.ShopService

	mockCreateShopReq  *entity.CreateShopRequest
	mockUpdateShopReq          *entity.UpdateShopRequest
	mockGetShopReq            *entity.GetShopRequest
	mockGetShopRes             entity.GetShopResponse
	mockGetShopEmptyShopRes entity.GetShopResponse
}

func (suite *ServiceList) SetupTest() {
	suite.mockShopRepo = new(mockPort.MockShopRepo)
	suite.service = NewShopService(suite.mockShopRepo)
	suite.mockCreateShopReq = &entity.CreateShopRequest{
		UserId:      "1",
		Name:        "Shop 1",
		Description: "Shop 1 Description",
		Terms:       "Shop 1 Terms",
	}
	suite.mockUpdateShopReq = &entity.UpdateShopRequest{
		UserId:      "1",
		Id:     	 "1",
		Name:        "Shop 1",
		Description: "Shop 1 Description",
		Terms:       "Shop 1 Terms",
	}
	suite.mockGetShopReq = &entity.GetShopRequest{
		Id:     "1",
	}
	suite.mockGetShopRes = entity.GetShopResponse{
		Id:          "1",
		Name:        "Shop 1",
		Description: "Shop 1 Description",
		Terms:       "Shop 1 Terms",
	}
	suite.mockGetShopEmptyShopRes = entity.GetShopResponse{}

}

// Testing CreateShop
func (u *ServiceList) TestCreateShop_Success() {
	ctx := context.Background()
	req := u.mockCreateShopReq
	u.mockShopRepo.Mock.On("IsUser", ctx, req).Return(true, nil)
	u.mockShopRepo.Mock.On("CreateShop", ctx, req).Return(mock.Anything, nil)
	_, err := u.service.CreateShop(ctx, req)

	u.Equal(nil, err)
}

func (u *ServiceList) TestCreateShop_UserNotFound() {
	ctx := context.Background()
	req := u.mockCreateShopReq
	errForbidden := errmsg.NewCustomErrors(403, errmsg.WithMessage("User not found"))

	u.mockShopRepo.Mock.On("IsUser", ctx, req).Return(false, nil)
	_, err := u.service.CreateShop(ctx, req)

	u.Equal(errForbidden, err)
}

func (u *ServiceList) TestCreateShop_IsUserError() {
	ctx := context.Background()
	req := u.mockCreateShopReq
	u.mockShopRepo.Mock.On("IsUser", ctx, req).Return(false, errors.New(mock.Anything))
	
	_, err := u.service.CreateShop(ctx, req)

	u.Equal(errors.New(mock.Anything), err)
}

func (u *ServiceList) TestCreateShop_CreateShopError() {
	ctx := context.Background()
	req := u.mockCreateShopReq
	u.mockShopRepo.Mock.On("IsUser", ctx, req).Return(true, nil)
	u.mockShopRepo.Mock.On("CreateShop", ctx, req).Return(nil, errors.New(mock.Anything))

	_, err := u.service.CreateShop(ctx, req)

	u.Equal(errors.New(mock.Anything), err)
}


func TestService(t *testing.T) {
	suite.Run(t, new(ServiceList))
}