package tg

import (
	"cms/internal/handler/telegohandlers"
	"context"
	"log"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type Router struct {
	Bot            *telego.Bot
	OrderHandler   *telegohandlers.OrderHandler
	ProductHandler *telegohandlers.ProductHandler
	UserHandler    *telegohandlers.UserHandler
}

func NewRouter(Oh *telegohandlers.OrderHandler, Ph *telegohandlers.ProductHandler, Uh *telegohandlers.UserHandler) *Router {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("env TOKEN not found")
	}

	admin := os.Getenv("ADMIN")
	if admin == "" {
		log.Fatal("env ADMIN not found")
	}

	bot, err := telego.NewBot(token, telego.WithDefaultLogger(false, true))
	if err != nil {
		log.Fatal(err)
	}

	return &Router{Bot: bot, OrderHandler: Oh, ProductHandler: Ph, UserHandler: Uh}
}

func (r *Router) Init() {
	updates, _ := r.Bot.UpdatesViaLongPolling(context.Background(), nil)
	bh, _ := th.NewBotHandler(r.Bot, updates)
	bh.Use(r.UserHandler.Middleware)
	r.RegisterHandlers(bh)
	bh.Start()
}

func (r *Router) RegisterHandlers(bh *th.BotHandler) {
	bh.Handle(r.UserHandler.HandleStart, th.CommandEqual("start"))
	bh.Handle(r.UserHandler.HandleName, r.UserHandler.StatePredicate("user_name"))
	bh.Handle(r.UserHandler.HandlePhone, r.UserHandler.StatePredicate("user_phone"))
	bh.Handle(r.UserHandler.HandleQuestionStart, th.TextEqual("❓Вопрос"))
	bh.Handle(r.UserHandler.HandleQuestionProcess, r.UserHandler.StatePredicate("user_question"))
	bh.Handle(r.OrderHandler.HandleStart, th.TextEqual("🛍Заказать товар"))
	bh.Handle(r.OrderHandler.HandleOrderProducts, r.UserHandler.StatePredicate("order_products"))
	bh.Handle(r.OrderHandler.HandleOrderProductsWeight, r.UserHandler.StatePredicate("order_products_weight"))
	bh.Handle(r.OrderHandler.HandleOrderProductsCount, r.UserHandler.StatePredicate("order_products_count"))
	bh.Handle(r.OrderHandler.HandleOrderDelivery, r.UserHandler.StatePredicate("order_delivery"))
	bh.Handle(r.OrderHandler.HandleOrderPayType, r.UserHandler.StatePredicate("order_paytype"))
	bh.Handle(r.OrderHandler.HandleOrderAddress, r.UserHandler.StatePredicate("order_address"))
	bh.Handle(r.ProductHandler.HandleStart, th.TextEqual("🛒Каталог"))
	bh.Handle(r.ProductHandler.HandleCatalogName, r.UserHandler.StatePredicate("product_list"))
	bh.Handle(r.ProductHandler.HandleCatalogWeight, r.UserHandler.StatePredicate("product_list_weight"))
	bh.Handle(r.ProductHandler.HandleCatalogAddName, r.UserHandler.StatePredicate("product_name"))
	bh.Handle(r.ProductHandler.HandleCatalogAddWeight, r.UserHandler.StatePredicate("product_weight"))
	bh.Handle(r.ProductHandler.HandleCatalogAddDescription, r.UserHandler.StatePredicate("product_description"))
	bh.Handle(r.ProductHandler.HandleCatalogAddImage, r.UserHandler.StatePredicate("product_image"))
	bh.Handle(r.ProductHandler.HandleCatalogAddPrice, r.UserHandler.StatePredicate("product_price"))
}
