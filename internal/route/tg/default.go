package tg

import (
	"cms/internal/handler/telegram"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Router struct {
	Bot             *tele.Bot
	OrderHandler    *telegram.OrderHandler
	ProductHandler  *telegram.ProductHandler
	UserHandler     *telegram.UserHandler
	QuestionHandler *telegram.QuestionHandler
}

func NewRouter(Oh *telegram.OrderHandler, Ph *telegram.ProductHandler, Uh *telegram.UserHandler, Qh *telegram.QuestionHandler) *Router {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("env TOKEN not found")
	}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Router{Bot: b, OrderHandler: Oh, ProductHandler: Ph, UserHandler: Uh, QuestionHandler: Qh}
}

func (r *Router) Init() {
	OrderGroup := r.Bot.Group()
	ProductGroup := r.Bot.Group()
	UserGroup := r.Bot.Group()
	QuestionGroup := r.Bot.Group()

	OrderGroup.Use(r.OrderHandler.Middleware)
	ProductGroup.Use(r.ProductHandler.Middleware)
	UserGroup.Use(r.UserHandler.Middleware)
	QuestionGroup.Use(r.QuestionHandler.Middleware)

	UserGroup.Handle("/start", r.UserHandler.Start)
	ProductGroup.Handle("🛒Каталог", r.ProductHandler.List)
	OrderGroup.Handle("🛍Заказать товар", r.OrderHandler.Start)
	QuestionGroup.Handle("❓Вопрос", r.QuestionHandler.Start)
	UserGroup.Handle("📞Контакты", r.UserHandler.Contact)

	r.Bot.Start()
}
