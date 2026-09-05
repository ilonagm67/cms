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

func (handler *OrderHandler) Middleware(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		err := next(c)
		return err
	}
}

func (handler *OrderHandler) Start(c tele.Context) error {
	return c.Send("Enter name: ")
}

func (handler *OrderHandler) Processing(c tele.Context) error {

}
