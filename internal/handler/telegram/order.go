package telegram

import (
	"cms/internal/service"

	tele "gopkg.in/telebot.v4"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(OrderService *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: OrderService}
}

func (o *OrderHandler) Start(c tele.Context) error {
	return nil
}
