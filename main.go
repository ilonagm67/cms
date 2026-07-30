package main

import (
	"cms/internal/repository/memory"
	"cms/internal/service"
)

func main() {
	userrepo := memory.NewUserRepository()
	orderrepo := memory.NewOrderRepository()
	productrepo := memory.NewProductRepository()

	service.NewUserService(userrepo)
	service.NewOrderService(orderrepo)
	service.NewProductService(productrepo)
}
