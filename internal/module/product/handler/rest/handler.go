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
	router.Post("/product", middleware.AuthBearer, h.CreateProduct)
	router.Get("/product/:id", h.GetDetailProduct)
	router.Patch("/product/:id", middleware.AuthBearer, h.UpdateProduct)
	router.Delete("/product/:id", middleware.AuthBearer, h.DeleteProduct)
	router.Get("/product", middleware.AuthBearer, h.GetProducts)
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.CreateProductRequest)
		ctx = c.Context()
		v = adapter.Adapters.Validator
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateProduct - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	// req.UserId = c.Params("user_id")

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

	return c.Status(fiber.StatusCreated).JSON(response.Success(resp, "Produk berhasil dibuat"))
}

func (h *productHandler) GetDetailProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.GetProductDetailRequest)
		ctx = c.Context()
		v = adapter.Adapters.Validator
	)

	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetProductDetail - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetDetailProduct(ctx, req)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("handler::GetDetailProduct - Failed to get product detail")
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}


	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h* productHandler) UpdateProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.UpdateProductRequest)
		ctx = c.Context()
		v = adapter.Adapters.Validator
		l = middleware.GetLocals(c)
	)

	req.Id = c.Params("id")

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::UpdateProduct - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.UserId = l.UserId

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::UpdateProduct - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.UpdateProduct(ctx, req)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("handler::UpdateProduct - Failed to update product")
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, "Produk berhasil diupdate"))
}

func (h* productHandler) DeleteProduct(c* fiber.Ctx) error{
	var(
		req = new(entity.DeleteProductRequest)
		ctx = c.Context()
		v = adapter.Adapters.Validator
		l = middleware.GetLocals(c)
	)

	req.Id = c.Params("id")
	req.UserId = l.UserId

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::DeleteProduct - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.DeleteProduct(ctx, req)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("handler::DeleteProduct - Failed to delete product")
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, "Produk berhasil dihapus"))
}

func (h* productHandler) GetProducts(c* fiber.Ctx) error{
	var(
		req = new(entity.GetProductsRequest)
		ctx = c.Context()
		v = adapter.Adapters.Validator
		l = middleware.GetLocals(c)
	)

	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::GetProducts - Parse request query")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.UserId = l.UserId
	req.SetDefault()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetProducts - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetProducts(ctx, req)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("handler::GetProducts - Failed to get products")
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}