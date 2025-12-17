package handlers

import (
	"database/sql"
	"errors"
	"example/shop-progect/internal/http/validator/dto"
	"example/shop-progect/internal/model"
	"example/shop-progect/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	product *service.ProductService
}

func NewProductHandler(product *service.ProductService) *ProductHandler {
	return &ProductHandler{product: product}
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	query := new(dto.ProductQuery)

	if err := c.Bind(query); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if err := c.Validate(query); err != nil {
		return err
	}

	products, err := h.product.GetListProduct(query.SKU, query.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByUUID(c echo.Context) error {
	id := c.Param("uuid")

	product, err := h.product.GetProductByUUID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DelProductByUUID(c echo.Context) error {
	id := c.Param("uuid")

	err := h.product.DeleteProduct(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "product not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ProductHandler) UpdateProductByUUID(c echo.Context) error {
	id := c.Param("uuid")

	req := new(dto.UpdateProduct)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	err := h.product.UpdateProduct(
		id,
		req.SKU,
		req.Name,
		req.Stock,
		req.ProductTypeID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "product not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.String(http.StatusOK, "OK")
}

func (h *ProductHandler) DelMultipleProducts(c echo.Context) error {
	req := new(dto.DelMultipleProducts)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	deleted, err := h.product.DeleteMultipleProducts(req.IDs)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"deleted": deleted,
	})
}

func (h *ProductHandler) AddProduct(c echo.Context) error {
	req := new(dto.AddProduct)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	user := c.Get("authUser").(*model.UserSess)

	err := h.product.CreateProduct(req.SKU, req.Name, req.Stock, req.ProductTypeID, user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"product_error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, "product add")
}
