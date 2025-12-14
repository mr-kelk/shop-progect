package handlers

import (
	"example/shop-progect/internal/http/validator/dto"
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
	h.product.GetListProduct()
	return c.JSON(http.StatusOK, "product list")
}

func (*ProductHandler) GetProductByUUID(c echo.Context) error {
	uuid := c.Param("uuid")
	return c.String(http.StatusOK, "product "+uuid)
}

func (h *ProductHandler) AddProduct(c echo.Context) error {
	req := new(dto.AddProduct)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"validation_error": err.Error(),
		})
	}

	err := h.product.CreateProduct()

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"product_error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, "product add")
}
