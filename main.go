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
	fsmrepo := memory.NewFSMRepository()
	orderrepo := memory.NewOrderRepository()
	productrepo := memory.NewProductRepository()
	userrepo := memory.NewUserRepository()

	fsmservice := service.NewFSMService(fsmrepo, orderrepo, productrepo, userrepo)
	orderservice := service.NewOrderService(orderrepo)
	productservice := service.NewProductService(productrepo)
	userservice := service.NewUserService(userrepo, fsmrepo)

	weborderhandler := web.NewOrderHandler(orderservice)
	webproducthandler := web.NewProductHandler(productservice)
	webuserhandler := web.NewUserHandler(userservice)

	tgorderhandler := telegram.NewOrderHandler(fsmservice)
	tgproducthandler := telegram.NewProductHandler(fsmservice)
	tguserhandler := telegram.NewUserHandler(fsmservice)

	httproute := http.NewRouter(weborderhandler, webproducthandler, webuserhandler)
	tgroute := tg.NewRouter(tgorderhandler, tgproducthandler, tguserhandler)

	go httproute.Init()
	go tgroute.Init()
	select {}
}
