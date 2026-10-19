package main

import (
	"cms/internal/handler/telegohandlers"
	"cms/internal/handler/web"
	"cms/internal/repository/memory"
	"cms/internal/repository/sqlite"
	"cms/internal/route/http"
	"cms/internal/route/tg"
	"cms/internal/service"
	"flag"
)

var (
	verboseFlag = flag.Bool("verbose", false, "verbose output")
	sqliteFlag  = flag.Bool("sqlite", false, "Use sqlite database")
	memoryFlag  = flag.Bool("memory", false, "Use memory database")
)

func main() {
	flag.Parse()
	if *verboseFlag {
		if *memoryFlag {
			fsmrepo := memory.NewFSMRepository()
			orderrepo := memory.NewOrderRepository()
			productrepo := memory.NewProductRepository()
			userrepo := memory.NewUserRepository()

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
			tgroute := tg.NewRouter(tgorderhandler, tgproducthandler, tguserhandler, true)

			go httproute.Init()
			go tgroute.Init()
			select {}
		} else if *sqliteFlag {
			fsmrepo := sqlite.NewFSMRepository()
			orderrepo := sqlite.NewOrderRepository()
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
			tgroute := tg.NewRouter(tgorderhandler, tgproducthandler, tguserhandler, true)

			go httproute.Init()
			go tgroute.Init()
			select {}
		} else {
			flag.Usage()
		}
	} else {
		if *memoryFlag {
			fsmrepo := memory.NewFSMRepository()
			orderrepo := memory.NewOrderRepository()
			productrepo := memory.NewProductRepository()
			userrepo := memory.NewUserRepository()

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
			tgroute := tg.NewRouter(tgorderhandler, tgproducthandler, tguserhandler, false)

			go httproute.Init()
			go tgroute.Init()
			select {}
		} else if *sqliteFlag {
			fsmrepo := sqlite.NewFSMRepository()
			orderrepo := sqlite.NewOrderRepository()
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
			tgroute := tg.NewRouter(tgorderhandler, tgproducthandler, tguserhandler, false)

			go httproute.Init()
			go tgroute.Init()
			select {}
		} else {
			flag.Usage()
		}
	}
}
