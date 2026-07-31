package telegram

import (
	"cms/internal/service"
	tele "gopkg.in/telebot.v4"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(Service *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: Service}
}
