package handler

import (
	"codebase-app/internal/adapter"
	"codebase-app/internal/middleware"
	"codebase-app/internal/module/product/entity"
	"codebase-app/internal/module/product/ports"
	"codebase-app/internal/module/product/repository"
	"codebase-app/internal/module/product/service"
	"codebase-app/pkg/errmsg"
	"codebase-app/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type productHandler struct {	
	service ports.ProductService
}

func NewProductHandler() *productHandler {
	var(
		handler = new(productHandler)
		repo = repository.NewProductRepository(adapter.Adapters.ShopeefunPostgres)
		service = service.NewProductService(repo)
	)
	handler.service = service

	return handler
}

func (h *productHandler) Register(router fiber.Router) {
	router.Post("/product", middleware.UserIdHeader, h.CreateProduct)
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.CreateProductRequest)
		ctx = c.Context()
		v = adapter.Adapters.Validator
		l = middleware.GetLocals(c)
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateProduct - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.UserId = l.UserId

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::CreateProduct - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.CreateProduct(ctx, req)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("handler::CreateProduct - Failed to create product")
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(resp, ""))
}