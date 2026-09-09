package telegram

import (
	"cms/internal/service"
)

type OrderHandler struct {
	OrderService *service.OrderService
	UserService  *service.UserService
}

func NewOrderHandler(Os *service.OrderService, Us *service.UserService) *OrderHandler {
	return &OrderHandler{OrderService: Os, UserService: Us}
}
