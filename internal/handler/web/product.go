package web

import (
	"cms/internal/service"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(Service *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: Service}
}
