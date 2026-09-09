package telegram

import (
	"cms/internal/service"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(ProductService *service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: ProductService}
}
