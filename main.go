package main

import (
	"cms/internal/handler/web"
	"cms/internal/repository/memory"
	"cms/internal/route/api"
	"cms/internal/service"
)

func main() {
	orderrepo := memory.NewOrderRepository()
	productrepo := memory.NewProductRepository()
	userrepo := memory.NewUserRepository()

	orderservice := service.NewOrderService(orderrepo)
	productservice := service.NewProductService(productrepo)
	userservice := service.NewUserService(userrepo)

	orderhandler := web.NewOrderHandler(orderservice)
	producthandler := web.NewProductHandler(productservice)
	userhandler := web.NewUserHandler(userservice)

	api.NewRouter(orderhandler, producthandler, userhandler)
}
