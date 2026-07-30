package main

import (
	"cms/internal/repository/memory"
	"cms/internal/handler/telegram"
	"cms/internal/route/tg"
	"cms/internal/service"
)

func main() {
	orderrepo := memory.NewOrderRepository()
	productrepo := memory.NewProductRepository()
	userrepo := memory.NewUserRepository()

	orderservice := service.NewOrderService(orderrepo)
	productservice := service.NewProductService(productrepo)
	userservice := service.NewUserService(userrepo)

	orderhandler := telegram.NewOrderHandler(orderservice)
	producthandler := telegram.NewProductHandler(productservice)
	userhandler := telegram.NewUserHandler(userservice)

	tg.NewRouter(orderhandler,producthandler,userhandler)
}
