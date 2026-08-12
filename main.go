package main

import (
	"cms/internal/handler/telegram"
	"cms/internal/handler/web"
	"cms/internal/repository/memory"
	"cms/internal/route/http"
	"cms/internal/route/tg"
	"cms/internal/service"
)

func main() {
	memory.NewStateRepository()
	orderrepo := memory.NewOrderRepository()
	productrepo := memory.NewProductRepository()
	userrepo := memory.NewUserRepository()

	orderservice := service.NewOrderService(orderrepo)
	productservice := service.NewProductService(productrepo)
	userservice := service.NewUserService(userrepo)

	weborderhandler := web.NewOrderHandler(orderservice)
	webproducthandler := web.NewProductHandler(productservice)
	webuserhandler := web.NewUserHandler(userservice)

	tgorderhandler := telegram.NewOrderHandler(orderservice)
	tgproducthandler := telegram.NewProductHandler(productservice)
	tguserhandler := telegram.NewUserHandler(userservice)
	tgquestionhandler := telegram.NewQuestionHandler(userservice)

	httproute := http.NewRouter(weborderhandler, webproducthandler, webuserhandler)
	tgroute := tg.NewRouter(tgorderhandler, tgproducthandler, tguserhandler, tgquestionhandler)

	go httproute.Init()
	go tgroute.Init()
	select {}
}
