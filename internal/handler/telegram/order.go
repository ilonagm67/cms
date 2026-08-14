package telegram

import (
	"cms/internal/service"

	tele "gopkg.in/telebot.v4"
)

type OrderHandler struct {
	OrderService *service.OrderService
	UserService  *service.UserService
}

func NewOrderHandler(Os *service.OrderService, Us *service.UserService) *OrderHandler {
	return &OrderHandler{OrderService: Os, UserService: Us}
}

func (o *OrderHandler) Start(c tele.Context) error {
	sender := c.Sender()
	o.OrderService.Add(sender.ID, sender.FirstName)
	return nil
}

func (o *OrderHandler) Hello() error {
	return nil
}
