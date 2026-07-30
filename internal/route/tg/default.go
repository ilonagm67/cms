package tg

import (
	"cms/internal/handler/telegram"
)

type Router struct {
	OrderHandler   *telegram.OrderHandler
	ProductHandler *telegram.ProductHandler
	UserHandler    *telegram.UserHandler
}

func NewRouter(OrderHandler *telegram.OrderHandler, ProductHandler *telegram.ProductHandler, UserHandler *telegram.UserHandler) *Router {
	return &Router{OrderHandler: OrderHandler, ProductHandler: ProductHandler, UserHandler: UserHandler}
}
