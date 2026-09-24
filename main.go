package main

import (
	"cms/internal/handler/telegohandlers"
	"cms/internal/handler/web"
	"cms/internal/repository/memory"
	"cms/internal/repository/sqlite"
	"cms/internal/route/http"
	"cms/internal/route/tg"
	"cms/internal/service"
)

func main() {
	fsmrepo := memory.NewFSMRepository()
	orderrepo := memory.NewOrderRepository()
	productrepo := sqlite.NewProductRepository()
	userrepo := sqlite.NewUserRepository()

	fsmservice := service.NewFSMService(fsmrepo, orderrepo, productrepo, userrepo)
	orderservice := service.NewOrderService(orderrepo)
	productservice := service.NewProductService(productrepo)
	userservice := service.NewUserService(userrepo)

	weborderhandler := web.NewOrderHandler(orderservice)
	webproducthandler := web.NewProductHandler(productservice)
	webuserhandler := web.NewUserHandler(userservice)

	tgorderhandler := telegohandlers.NewOrderHandler(fsmservice, productservice, orderservice)
	tgproducthandler := telegohandlers.NewProductHandler(fsmservice, productservice)
	tguserhandler := telegohandlers.NewUserHandler(fsmservice, userservice)

	httproute := http.NewRouter(weborderhandler, webproducthandler, webuserhandler)
	tgroute := tg.NewRouter(tgorderhandler, tgproducthandler, tguserhandler)

	go httproute.Init()
	go tgroute.Init()
	select {}
}
