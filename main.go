package main

import (
	"cms/internal/handler/telegram"
	"cms/internal/handler/web"
	"cms/internal/repository/memory"
	"cms/internal/route/api"
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

	apiorderhandler := web.NewOrderHandler(orderservice)
	apiproducthandler := web.NewProductHandler(productservice)
	apiuserhandler := web.NewUserHandler(userservice)

	tgorderhandler := telegram.NewOrderHandler(orderservice)
	tgproducthandler := telegram.NewProductHandler(productservice)
	tguserhandler := telegram.NewUserHandler(userservice)
	tgquestionhandler := telegram.NewQuestionHandler(userservice)

	go api.NewRouter(apiorderhandler, apiproducthandler, apiuserhandler)
	go tg.NewRouter(tgorderhandler, tgproducthandler, tguserhandler, tgquestionhandler)
	select {}
}
